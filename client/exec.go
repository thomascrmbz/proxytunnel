package client

import (
	"thomascrmbz.com/proxytunnel"
)

func (pc *ProxyClient) Execute(agentID int, args ...string) {
	pc.execTunnelCmd(proxytunnel.Exec, agentID, args...)
}
