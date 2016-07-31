package slack

import (
	"log"

	"github.com/ebenoist/ply"
)

func init() {
	ply.RegisterTaskType("Slack", New)
}

func New(name string, options map[string]string) ply.Task {
	return &Slack{Name: name, options: options}
}

type Slack struct {
	Name    string
	Options map[string]string
}

func (t *Slack) Name() string {
	return t.Name
}

func (t *Slack) Options() map[string]string {
	return t.Options
}

func (t *Slack) Run(host string) error {
	log.Printf("slacking on %s", host)
}
