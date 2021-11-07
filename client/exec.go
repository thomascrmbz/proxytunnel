package client

import (
	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
)

func (pc *ProxyClient) Execute(agentServer agent.Agent, args ...string) {
	pc.execTunnelCmd(proxytunnel.Exec, agentServer, args...)
}
