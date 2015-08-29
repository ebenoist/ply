package main

import "testing"

func assertEq(t *testing.T, exp, act string) {
	if exp != act {
		t.Errorf("Expected %s to equal %s", exp, act)
	}
}

func Test_LoadConfigUnmarshallsTheYAML(t *testing.T) {
	yaml := `
Version: 1
Vars:
  AppName: my-app
Tasks:
  Deploy:
    - docker run foo
DeployEnvs:
  staging:
    Hosts:
      - foo-service.local
`
	cfg := LoadConfig([]byte(yaml), Vars{}, "staging")

	assertEq(t, cfg.Version, "1")
	assertEq(t, cfg.Tasks["Deploy"][0], "docker run foo")
	assertEq(t, cfg.DeployEnvs["staging"].Hosts[0], "foo-service.local")
}

func Test_TheConfigCanBeTemplatedByVarsPassedIn(t *testing.T) {
	yaml := `
Tasks:
  Deploy:
    - docker run {{.DockerTag}}
`
	vars := Vars{
		"DockerTag": "foo-tag",
	}

	cfg := LoadConfig([]byte(yaml), vars, "staging")
	assertEq(t, cfg.Tasks["Deploy"][0], "docker run foo-tag")
}

func Test_TheConfigCanBeTemplatedByEnvVars(t *testing.T) {
	yaml := `
Vars:
  AppName: my-app
Tasks:
  Deploy:
    - docker run {{.AppName}} -e {{.JRubyOpts}}
DeployEnvs:
  production:
    Vars:
      JRubyOpts: "-J-Xmn1024"
`
	cfg := LoadConfig([]byte(yaml), Vars{}, "production")
	assertEq(t, cfg.Tasks["Deploy"][0], "docker run my-app -e -J-Xmn1024")
}

func Test_ConfigCanBeTemplatesByVars(t *testing.T) {
	yaml := `
Vars:
  AppName: my-app
Tasks:
  Deploy:
    - docker run {{.AppName}}
`

	cfg := LoadConfig([]byte(yaml), Vars{}, "production")
	assertEq(t, cfg.Tasks["Deploy"][0], "docker run my-app")
}
