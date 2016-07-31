package ply

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
)

var green = ansi.ColorFunc("green")
var red = ansi.ColorFunc("red")
var yellow = ansi.ColorFunc("yellow")
var logger = NewLogger()

func Ply(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("ERROR: Please specify a task\n\n")
		cli.ShowAppHelp(c)
		return
	}

	// task := c.Args()[0]

	// tplVars := c.GlobalString("vars")
	// cfgPath := c.GlobalString("config")
	// env := c.GlobalString("environment")
	// host := c.GlobalString("host")

	// vars := parseVars(tplVars)
	// file, err := ioutil.ReadFile(cfgPath)
	// if err != nil {
	// fmt.Printf("ERROR: could not find the config file: %s\n\n", cfgPath)
	// cli.ShowAppHelp(c)
	// return
	// }

	// cfg := LoadConfig(file, vars, env)
	// if host != "" {
	// Run(task, []string{host}, cfg)
	// } else {
	// Run(task, cfg.DeployEnvs[env].Hosts, cfg)
	// }
}

func parseVars(raw string) Vars {
	vars := Vars{}
	if raw == "" {
		return vars
	}

	for _, v := range strings.Split(raw, ",") {
		pair := strings.Split(v, "=")
		vars[pair[0]] = pair[1]
	}

	return vars
}
