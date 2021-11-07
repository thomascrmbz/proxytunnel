package proxytunnel

type TunnelCmd string

var (
	Exec  TunnelCmd = "TUNNEL_EXEC"
	Shell TunnelCmd = "TUNNEL_SHELL"
	Proxy TunnelCmd = "TUNNEL_PROXY"
)
