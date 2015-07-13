package main

import (
	"bytes"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version    string                  `yaml:"Version"`
	AppName    string                  `yaml:"AppName"`
	DeployUser string                  `yaml:"DeployUser"`
	Container  string                  `yaml:"Container"`
	BeforeRun  []string                `yaml:"BeforeRun"`
	RunCommand string                  `yaml:"RunCommand"`
	AfterRun   []string                `yaml:"AfterRun"`
	DeployEnvs map[string]DeployEnvCfg `yaml:"DeployEnvs"`
	Plugins    map[string]PluginCfg    `yaml:"Plugins"`
}

type DeployEnvCfg struct {
	Hosts []string `yaml:"Hosts"`
	Vars  Vars     `yaml:"Vars"`
}

type PluginCfg map[string]interface{}
type Vars map[string]string

func LoadConfig(yml []byte, vars Vars, env string) Config {
	c := parseConfig(parseTmpl(yml, vars))
	cfgVars := c.DeployEnvs[env].Vars

	for k, v := range cfgVars {
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
