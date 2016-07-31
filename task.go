package ply

var TaskTypes = map[string]func(string, interface{}) Task{}

type Task interface {
	GetName() string
	Run(string) error
}

func RegisterTaskType(name string, constructor func(string, interface{}) Task) {
	TaskTypes[name] = constructor
}
