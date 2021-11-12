package client

import (
	"thomascrmbz.com/proxytunnel"
)

func (pc *ProxyClient) Shell(agentID int) {
	pc.execTunnelCmd(proxytunnel.Shell, agentID)
}
