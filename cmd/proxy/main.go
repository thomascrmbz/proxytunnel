package main

import (
	"log"

	"thomascrmbz.com/proxytunnel/agent"
	proxytunnel "thomascrmbz.com/proxytunnel/proxy"
)

func main() {
	proxy := proxytunnel.Proxy{
		Port:        8020,
		AuthHandler: func() bool { return true },
	}

	proxy.AddAgent(
		agent.Agent{
			ID:   1,
			Name: "production",
			Port: 43023,
			IP:   "devops@staging.pulu.devbitapp.be",
		},
		agent.Agent{
			ID:   2,
			Name: "staging",
			Port: 43022,
			IP:   "devops@staging.pulu.devbitapp.be",
		},
		// agent.Agent{
		// 	ID:   3,
		// 	Name: "devboard",
		// },
	)

	log.Fatal(proxy.ListenAndServe())
}
