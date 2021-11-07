package client

import (
	"strconv"

	"thomascrmbz.com/proxytunnel"
	"thomascrmbz.com/proxytunnel/agent"
)

func (pc *ProxyClient) Proxy(agentServer agent.Agent, port int) {
	pc.execTunnelCmd(proxytunnel.Proxy, agentServer, strconv.Itoa(port))
}
