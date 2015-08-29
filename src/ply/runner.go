package main

import (
	"fmt"
	"log"
)

func Run(task string, hosts []string, config Config) {
	for _, host := range hosts {
		log.Printf("Running %s on %s", task, host)
		client, err := NewAgentClient(host, config.DeployUser)
		if err != nil {
			log.Fatalf("Could not connect %s", err)
		}

		log.SetPrefix(fmt.Sprintf("%s - ", host))
		RunTask(config.Tasks[task], client, config.Tasks)
	}
}

func RunTask(steps []string, client SSHClient, tasks Tasks) {
	for _, step := range steps {
		if tasks[step] != nil {
			RunTask(tasks[step], client, tasks)
		} else {
			log.Printf("Running: %s", step)
			err := client.Run(step)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
