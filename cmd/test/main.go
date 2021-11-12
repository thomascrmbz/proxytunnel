package main

import (
	proxytunnel "thomascrmbz.com/proxytunnel/client"
)

func main() {
	proxy := proxytunnel.NewProxyClient("pulu.trikthom.com", 8020)

	// proxy.Execute(aserver, "echo $(date) hostname of agent server: $(hostname)")
	proxy.Shell(2)
}
