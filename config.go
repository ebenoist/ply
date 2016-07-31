package ply

import (
	"bytes"
	"reflect"
	"text/template"

	"gopkg.in/yaml.v2"
)

var Tasks = map[string]Task{}
var Config PlyConfig

type Vars map[string]string

type PlyConfig struct {
	Version string                 `yaml:"version"`
	Vars    Vars                   `yaml:"vars"`
	Tasks   map[string]interface{} `yaml:"tasks"`
	Remote  map[string]interface{} `yaml:"remote"`
}

func LoadConfig(yml []byte, runtimeVars map[string]string, env string) {
	rawConfig := parseConfig(yml)
	vars := map[string]string{}

	cfgVars := rawConfig.Vars
	// envVars := rawConfig.DeployEnvs[env].Vars

	for k, v := range cfgVars {
		vars[k] = v
	}

	// for k, v := range envVars {
	// vars[k] = v
	// }

	for k, v := range runtimeVars {
		vars[k] = v
	}

	c := parseConfig(parseTmpl(yml, vars))

	Config = c
	Tasks = buildTasks(c)
}

func buildTasks(raw PlyConfig) map[string]Task {
	tasks := map[string]Task{}

	for name, options := range raw.Tasks {
		kind := reflect.TypeOf(options).Kind()
		var taskType string

		if kind == reflect.String || kind == reflect.Slice {
			taskType = "remote"
		}

		if kind == reflect.Map {
			o := options.(map[interface{}]interface{})
			taskType = o["runner"].(string)
		}

		if taskType == "" {
			panic("Bad config")
		}

		task := TaskTypes[taskType](name, options)
		tasks[name] = task
	}

	return tasks
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

func parseConfig(yml []byte) PlyConfig {
	c := PlyConfig{}
	err := yaml.Unmarshal(yml, &c)

	if err != nil {
		panic("yml: " + err.Error())
	}

	return c
}
