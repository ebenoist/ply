package main

import (
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func NewSession(url string, user string) (*SSHSession, error) {
	cfg := &ssh.ClientConfig{
		User: user,
		Auth: sshAgentAuth(),
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

func sshAgentAuth() ssh.AuthMethod {
	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	agent := agent.NewClient(sock)

	signers, err := agent.Signers()
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signers...)
}
