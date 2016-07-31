package remote

import (
	"bytes"
	"errors"
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
	cfg := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{SSHAgent()},
	}
	cfg.SetDefaults()

	url := fmt.Sprintf("%s:22", host)
	conn, err := ssh.Dial("tcp", url, cfg)

	if err != nil {
		return nil, err
	}

	return &AgentClient{conn: conn}, err
}

func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}

	return nil
}

func (c *AgentClient) Run(cmd string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var errs bytes.Buffer
	var infos bytes.Buffer

	session.Stdout = &infos
	session.Stderr = &errs

	err = session.Run(cmd)

	if len(errs.String()) > 0 {
		return errors.New(errs.String())
	}

	log.Printf("%s", errs.String())
	return err
}
