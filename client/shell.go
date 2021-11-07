package client

import (
	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
)

func (pc *ProxyClient) Shell(agentServer agent.Agent) {
	pc.execTunnelCmd(proxytunnel.Shell, agentServer)
}
