package proxy

import (
	"fmt"
	"strconv"

	"github.com/gliderlabs/ssh"
	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
	"thomascrmbz.com/proxytunnel/proxy/handler"
)

type Proxy struct {
	Port        int
	AuthHandler func() bool

	agents []*agent.Agent
}

func (p *Proxy) ListenAndServe() error {
	sshServer := ssh.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(p.Port),
		Handler: func(s ssh.Session) {
			handler.DefaultHandler(p.findAgent(s), s)
		},
		PublicKeyHandler: func(ctx ssh.Context, key ssh.PublicKey) bool {
			ctx.SetValue("sshPublicKey", key.Marshal())
			return true
		},
	}

	fmt.Println("Proxytunnel started on " + sshServer.Addr + " with " + strconv.Itoa(len(p.agents)) + " agents")

	return sshServer.ListenAndServe()
}

func (p *Proxy) AddAgent(a agent.Agent, agents ...agent.Agent) {
	p.agents = append(p.agents, &a)
	for _, ag := range agents {
		p.agents = append(p.agents, &ag)
	}
}

func (p *Proxy) findAgent(s ssh.Session) *agent.Agent {
	agentId, _ := strconv.Atoi(s.Command()[1])
	for _, a := range p.agents {
		if a.ID == agentId {
			return a
		}
	}
	s.Exit(int(proxytunnel.AGENT_NOT_FOUND))
	return &agent.Agent{}
}
