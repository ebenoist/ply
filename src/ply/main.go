package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

func main() {
	app := cli.NewApp()
	app.Name = "ply"
	app.Usage = "Deployment made simple"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "The path to your ply config",
		},
		cli.StringFlag{
			Name:  "environment, e",
			Usage: "The environment to execute the task",
		},
		cli.StringFlag{
			Name:  "vars, variables",
			Usage: "Comma separated list of variables (ie. AppName=MyApp,Revision=v1.2.3)",
		},
		cli.StringFlag{
			Name:  "host",
			Usage: "Host override",
		},
	}

	app.Action = Action
	app.Run(os.Args)
}

func Action(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("ERROR: Please specify a a task\n\n")
		cli.ShowAppHelp(c)
		return
	}

	task := c.Args()[0]

	tplVars := c.GlobalString("vars")
	cfgPath := c.GlobalString("config")
	env := c.GlobalString("environment")
	host := c.GlobalString("host")

	vars := parseVars(tplVars)
	file, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		fmt.Printf("ERROR: could not find the config file: %s\n\n", cfgPath)
		cli.ShowAppHelp(c)
		return
	}

	cfg := LoadConfig(file, vars, env)
	if host != "" {
		Run(task, []string{host}, cfg)
	} else {
		Run(task, cfg.DeployEnvs[env].Hosts, cfg)
	}
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
