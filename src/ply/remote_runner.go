package main

import "log"

// Define a set of tasks
// When I run a task, see if the task exists via the Name if so run it
type RemoteTask struct {
	Name    string
	Options map[string]string
}

func (t *RemoteTask) Name() string {
	return t.Name
}

func (t *RemoteTask) Options() map[string]string {
	return t.Options
}

func (t *RemoteTask) Run(host string) error {
	user := Config.Plugins.Remote.User
	client := NewAgentClient(host, user)

	log.Printf("%s", green(step))
	err := client.Run(step)
	if err != nil {
		log.Fatalf("%s", red(err.Error()))
	}
}
