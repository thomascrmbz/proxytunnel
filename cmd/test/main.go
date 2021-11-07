package main

import (
	"thomascrmbz.com/proxytunnel/agent"
	proxytunnel "thomascrmbz.com/proxytunnel/client"
)

func main() {
	proxy := proxytunnel.NewProxyClient("devops@pulu.trikthom.com", 22)
	aserver := agent.Agent{ID: 1, Name: "test-server-1"}

	proxy.Execute(aserver, "echo $(date) hostname of agent server: $(hostname)")
	proxy.Shell(aserver)
}
