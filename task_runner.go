package ply

import (
	"fmt"
	"log"
)

func RunTasks(hosts []string, tasks []Task) {
	for _, host := range hosts {
		for _, task := range tasks {
			msg := fmt.Sprintf("Running %s on %s", task.GetName(), host)
			log.Printf("%s", yellow(msg))
			prefix := fmt.Sprintf("%s ", yellow(host))
			log.SetPrefix(prefix)
			task.Run(host)
		}
	}
}
