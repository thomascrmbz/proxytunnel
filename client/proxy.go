package client

import (
	"strconv"

	"thomascrmbz.com/proxytunnel"
)

func (pc *ProxyClient) Proxy(agentID int, port int) {
	pc.execTunnelCmd(proxytunnel.Proxy, agentID, strconv.Itoa(port))
}
