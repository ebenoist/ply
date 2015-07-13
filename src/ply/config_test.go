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
AppName: my-app
DeployUser: user
Container: foo-container
BeforeRun:
  - echo 'hello'
RunCommand: docker run foo
AfterRun:
  - echo 'done'
  - echo 'still done'
DeployEnvs:
  staging:
    Hosts:
      - foo-service.local
`
	cfg := LoadConfig([]byte(yaml), Vars{}, "staging")

	tests := map[string]string{
		cfg.Version:    "1",
		cfg.AppName:    "my-app",
		cfg.DeployUser: "user",
		cfg.Container:  "foo-container",
		cfg.RunCommand: "docker run foo",
	}

	for k, v := range tests {
		assertEq(t, v, k)
	}

	assertEq(t, cfg.AfterRun[0], "echo 'done'")
	assertEq(t, cfg.AfterRun[1], "echo 'still done'")
	assertEq(t, cfg.DeployEnvs["staging"].Hosts[0], "foo-service.local")
}

func Test_TheConfigCanBeTemplatedByVarsPassedIn(t *testing.T) {
	yaml := `
AppName: my-app
RunCommand: docker run my-app {{.DockerTag}}
`
	vars := Vars{
		"DockerTag": "foo-tag",
	}

	cfg := LoadConfig([]byte(yaml), vars, "staging")
	assertEq(t, cfg.RunCommand, "docker run my-app foo-tag")
}

func Test_TheConfigCanBeTemplatedByStaticVars(t *testing.T) {
	yaml := `
AppName: my-app
RunCommand: "docker run my-app -e JRubyOpts=\"{{.JRubyOpts}}\""
DeployEnvs:
  production:
    Vars:
      JRubyOpts: -J-Xmn5376m -J-Xms7168m -J-Xmx7168m
`

	cfg := LoadConfig([]byte(yaml), Vars{}, "production")
	assertEq(t, cfg.RunCommand, "docker run my-app -e JRubyOpts=\"-J-Xmn5376m -J-Xms7168m -J-Xmx7168m\"")
}
