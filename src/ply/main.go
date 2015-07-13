package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	cfgPath := flag.String("c", "", "The path to the config")
	deployEnv := flag.String("e", "", "The deploy env")
	tplVars := flag.String("var", "", "Override variables")

	flag.Parse()

	vars := parseVars(*tplVars)

	file, err := ioutil.ReadFile(*cfgPath)

	if err != nil {
		msg := fmt.Sprintf("Could not find config file: %s", *cfgPath)
		panic(msg)
	}

	cfg := LoadConfig(file, vars, *deployEnv)
	for h := range cfg.DeployEnvs[deployEnv].Hosts {
	}
}

func parseVars(raw string) Vars {
	vars := Vars{}
	for _, v := range strings.Split(raw, ",") {
		pair := strings.Split(v, "=")
		vars[pair[0]] = pair[1]
	}

	return vars
}
