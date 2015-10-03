package main

type Plugin interface {
	Run(host string, client SSHClient, config Config) error
}
