package main

import (
	"fmt"
	"log"
)

func Run(task string, hosts []string, config Config) {
	for _, host := range hosts {
		msg := fmt.Sprintf("Running %s on %s", task, host)
		log.Printf("%s", yellow(msg))
		client, err := NewAgentClient(host, config.DeployUser)
		if err != nil {
			log.Fatalf("%s", red(err.Error()))
		}

		prefix := fmt.Sprintf("%s ", yellow(host))
		log.SetPrefix(prefix)
		RunTask(config.Tasks[task], client, config.Tasks)
	}
}

func RunTask(steps []string, client SSHClient, tasks Tasks) {
	for _, step := range steps {
		if tasks[step] != nil {
			RunTask(tasks[step], client, tasks)
		} else {
			log.Printf("%s", green(step))
			err := client.Run(step)
			if err != nil {
				log.Fatalf("%s", red(err.Error()))
			}
		}
	}
}
