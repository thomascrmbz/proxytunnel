package handler

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
)

func DefaultHandler(a *agent.Agent, s ssh.Session) {
	command := s.Command()

	if len(command) >= 2 {
		fmt.Println(command)

		switch command[0] {
		case string(proxytunnel.Shell):
			ShellHandler(a, s)
		case string(proxytunnel.Exec):
			ExecHandler(a, s)
		case string(proxytunnel.Proxy):
			ProxyHandler(a, s)
		default:
			s.Exit(int(proxytunnel.COMMAND_NOT_FOUND))
		}
	} else {
		s.Exit(int(proxytunnel.NOT_ALLOWED))
	}
}
