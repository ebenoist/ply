package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type SSHClient interface {
	Run(cmd string) error
}

type AgentClient struct {
	conn *ssh.Client
}

func NewAgentClient(host string, user string) (*AgentClient, error) {
	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	agent := agent.NewClient(sock)

	signers, err := agent.Signers()
	if err != nil {
		return nil, err
	}

	cfg := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signers...)},
	}
	cfg.SetDefaults()

	url := fmt.Sprintf("%s:22", host)
	conn, err := ssh.Dial("tcp", url, cfg)

	if err != nil {
		return nil, err
	}

	return &AgentClient{conn: conn}, err
}

func (c *AgentClient) Run(cmd string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	err = session.Run(cmd)
	log.Printf("Response: %s", b.String())
	return err
}
