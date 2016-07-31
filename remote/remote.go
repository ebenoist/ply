package remote

import (
	"reflect"

	"github.com/ebenoist/ply"
)

func New(name string, cmd interface{}) ply.Task {
	var cmds []string

	if reflect.TypeOf(cmd).Kind() == reflect.String {
		cmds = append(cmds, cmd.(string))
	}

	if reflect.TypeOf(cmd).Kind() == reflect.Slice {
		for _, c := range cmd.([]interface{}) {
			cmds = append(cmds, c.(string))
		}
	}

	return &Task{
		Name: name,
		Cmds: cmds,
	}
}

type Task struct {
	Name string
	Cmds []string
}

func (t *Task) GetName() string {
	return t.Name
}

func (t *Task) Run(host string) error {
	user := ply.Config.Remote["user"].(string)
	client, _ := NewAgentClient(host, user)

	for _, c := range t.Cmds {
		sub, ok := ply.Tasks[c]
		if ok {
			sub.Run(host)
		} else {
			client.Run(c)
		}
	}

	return nil
}
