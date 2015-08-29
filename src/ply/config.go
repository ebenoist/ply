package main

import (
	"bytes"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Tasks map[string][]string
type Vars map[string]string

type Config struct {
	Version    string               `yaml:"Version"`
	Vars       Vars                 `yaml:"Vars"`
	Tasks      Tasks                `yaml:"Tasks"`
	DeployUser string               `yaml:"DeployUser"`
	DeployEnvs map[string]DeployEnv `yaml:"DeployEnvs"`
}

type DeployEnv struct {
	Hosts []string `yaml:"Hosts"`
	Vars  Vars     `yaml:"Vars"`
}

func LoadConfig(yml []byte, vars Vars, env string) Config {
	c := parseConfig(parseTmpl(yml, vars))

	cfgVars := c.Vars
	envVars := c.DeployEnvs[env].Vars

	for k, v := range cfgVars {
		vars[k] = v
	}

	for k, v := range envVars {
		vars[k] = v
	}

	c = parseConfig(parseTmpl(yml, vars))
	return c
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

func parseConfig(yml []byte) Config {
	c := Config{}
	err := yaml.Unmarshal(yml, &c)

	if err != nil {
		panic("yml: " + err.Error())
	}

	return c
}
