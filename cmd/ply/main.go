package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ebenoist/ply"
	"github.com/ebenoist/ply/remote"
)

func main() {
	ply.RegisterTaskType("remote", remote.New)

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

	app.Action = ply.Ply
	app.Run(os.Args)
}
