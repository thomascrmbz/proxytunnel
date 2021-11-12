package main

import (
	proxytunnel "thomascrmbz.com/proxytunnel/client"
)

func main() {
	proxy := proxytunnel.NewProxyClient("pulu.trikthom.com", 8020)

	// proxy.Execute(2, "date")
	proxy.Shell(1)
}
