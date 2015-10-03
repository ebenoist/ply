package main

import (
	"bytes"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Vars map[string]string
type RawConfig struct {
	Version     string                 `yaml:"Version"`
	Vars        Vars                   `yaml:"Vars"`
	Tasks       map[string]interface{} `yaml:"Tasks"`
	Environment map[string]interface{} `yaml:"Environment"`
	Plugins     map[string]interface{} `yaml:"Plugins"`
}

type Config struct {
	Tasks       []*Task
	Environment Environment
}

type Environment struct {
	Hosts []string
	User  string
}

func NewConfig(yml []byte, runtimeVars map[string]string, env string) Config {
	rawConfig := parseConfig(yml)

	cfgVars := rawConfig.Vars
	envVars := rawConfig.DeployEnvs[env].Vars

	for k, v := range cfgVars {
		vars[k] = v
	}

	for k, v := range envVars {
		vars[k] = v
	}

	for k, v := range runtimeVars {
		vars[k] = v
	}

	c = parseConfig(parseTmpl(yml, vars))
	return &Config{
		Tasks:       buildTasks(c),
		Environment: buildEnv(c),
	}
}

func buildEnv(raw RawConfig) Environment {
}

func buildTasks(raw RawConfig) []*Task {

}

func parseTmpl(tpl []byte, vars Vars) []byte {
	t, err := template.New("config").Parse(string(tpl))

	buff := &bytes.Buffer{}
	t.Execute(buff, vars)

	if err != nil {
		panic("template: " + err.Error())
	}

	return buff.Bytes()
}

func parseConfig(yml []byte) RawConfig {
	c := RawConfig{}
	err := yaml.Unmarshal(yml, &c)

	if err != nil {
		panic("yml: " + err.Error())
	}

	return c
}
