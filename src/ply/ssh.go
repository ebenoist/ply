package main

import (
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func NewSession(url string, user string) (*ssh.Session, error) {
	auths := make([]ssh.AuthMethod, 1)
	auth, _ := sshAgentAuth()
	auths = append(auths, auth)

	cfg := &ssh.ClientConfig{
		User: user,
		Auth: auths,
	}
	cfg.SetDefaults()

	client, err := ssh.Dial("tcp", url, cfg)

	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func sshAgentAuth() (ssh.AuthMethod, error) {
	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	agent := agent.NewClient(sock)

	signers, err := agent.Signers()
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signers...), nil
}
