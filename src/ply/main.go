package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {
	host := flag.String("host", "", "The host to connect to")
	port := flag.Int("port", 22, "The port to connect to")
	flag.Parse()

	connect(*host, *port)
}

func connect(host string, port int) {
	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Fatal(err)
	}

	agent := agent.NewClient(sock)

	signers, err := agent.Signers()
	if err != nil {
		log.Fatal(err)
	}

	auths := []ssh.AuthMethod{ssh.PublicKeys(signers...)}

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &ssh.ClientConfig{
		User: u.Username,
		Auth: auths,
	}
	cfg.SetDefaults()

	url := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", url, cfg)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("We've got a live session!")
}
