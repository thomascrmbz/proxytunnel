package handler

import (
	"github.com/gliderlabs/ssh"
	"thomascrmbz.com/proxytunnel/agent"
)

func ExecHandler(a *agent.Agent, s ssh.Session) {
	sshExe(a, s, sshOptions{
		CopyStdin:  true,
		CopyStdout: true,
	}, s.Command()[2:]...)
}
