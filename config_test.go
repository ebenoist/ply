package ply_test

import (
	"testing"

	"github.com/ebenoist/ply"
)

type TestTask struct {
	Name    string
	Options map[string]string
}

func (t *TestTask) GetName() string {
	return t.Name
}

func (t *TestTask) Run(host string) error {
	return nil
}

func NewTestTask(name string, options interface{}) ply.Task {
	return &TestTask{
		Name:    name,
		Options: options.(map[string]string),
	}
}

func Setup() {
	ply.RegisterTaskType("test", NewTestTask)
}

func assertEq(t *testing.T, exp, act string) {
	if exp != act {
		t.Errorf("Expected %s to equal %s", exp, act)
	}
}

func Test_LoadConfigUnmarshallsTheYAML(t *testing.T) {
	ply.RegisterTaskType("test", NewTestTask)

	yaml := `
version: 2
vars:
  app_name: my-app
tasks:
  deploy:
    runner: test
    cmd: docker run foo
  start:
    - deploy
    - touch /tmp/foo
remote:
  staging:
    user: erik
    hosts:
      - foo-service.local
      - bar-service.local
`
	ply.LoadConfig([]byte(yaml), ply.Vars{}, "staging")
	if ply.Tasks["deploy"].GetName() != "deploy" {
		t.Error("Expected name to be deploy")
	}

	if ply.Tasks["start"].GetName() != "start" {
		t.Error("Expected name to be start")
	}
}

func Test_TheConfigCanBeTemplatedByVarsPassedIn(t *testing.T) {
	yaml := `
tasks:
  deploy:
    - docker run {{.DockerTag}}
`
	vars := ply.Vars{
		"DockerTag": "foo-tag",
	}

	ply.LoadConfig([]byte(yaml), vars, "staging")
	assertEq(t, ply.Tasks["Deploy"].GetName(), "docker run foo-tag")
}

// func Test_TheConfigCanBeTemplatedByEnvVars(t *testing.T) {
// yaml := `
// Vars:
// AppName: my-app
// Tasks:
// Deploy:
// - docker run {{.AppName}} -e {{.JRubyOpts}}
// DeployEnvs:
// production:
// Vars:
// JRubyOpts: "-J-Xmn1024"
// `
// cfg := NewConfig([]byte(yaml), Vars{}, "production")
// assertEq(t, cfg.Tasks["Deploy"][0], "docker run my-app -e -J-Xmn1024")
// }

// func Test_ConfigCanBeTemplatesByVars(t *testing.T) {
// yaml := `
// Vars:
// AppName: my-app
// Tasks:
// Deploy:
// - docker run {{.AppName}}
// `

// cfg := NewConfig([]byte(yaml), Vars{}, "production")
// assertEq(t, cfg.Tasks["Deploy"][0], "docker run my-app")
// }
