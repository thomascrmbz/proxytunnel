package client

import (
	"thomascrmbz.com/proxytunnel"
)

func (pc *ProxyClient) Proxy(agentID int, port string) {
	pc.execTunnelCmd(proxytunnel.Proxy, agentID, port)
}
