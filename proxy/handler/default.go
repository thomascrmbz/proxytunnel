package handler

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
)

func DefaultHandler(a *agent.Agent, s ssh.Session) {
	command := s.Command()
	fmt.Println(command)

	// Check permissions from ssh public key
	// sshPublicKey := s.Context().Value("sshPublicKey")

	switch command[0] {
	case string(proxytunnel.Shell):
		ShellHandler(a, s)
	default:
		s.Exit(int(proxytunnel.OK))
	}
}
