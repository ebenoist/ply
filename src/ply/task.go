package main

type Task interface {
	Name() string
	Options() map[string]string
	Run(string) error
}
