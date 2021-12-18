package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh"
	"golang.org/x/oauth2"
	"thomascrmbz.com/proxytunnel/agent"
	proxytunnel "thomascrmbz.com/proxytunnel/proxy"
)

func main() {
	proxy := proxytunnel.Proxy{
		Port:        8020,
		AuthHandler: auth,
	}

	proxy.AddAgent(
		&agent.Agent{
			ID:   1,
			Name: "production",
			Port: 43023,
			IP:   "devops@staging.pulu.devbitapp.be",
		},
		&agent.Agent{
			ID:   2,
			Name: "staging",
			Port: 43022,
			IP:   "devops@staging.pulu.devbitapp.be",
		},
		&agent.Agent{
			ID:   3,
			Name: "devboard",
			Port: 43025,
			IP:   "ubuntu@localhost",
		},
	)

	log.Fatal(proxy.ListenAndServe())
}

var client *github.Client
var ctx = context.Background()

func init() {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}

func auth(key ssh.PublicKey) bool {
	for _, ghKey := range getKeys(getOrgUsers()) {
		if bytes.Equal(ghKey.Marshal(), key.Marshal()) {
			return true
		}
	}
	return false
}

func getOrgUsers() []*github.User {
	users, _, _ := client.Organizations.ListMembers(ctx, "vives-projectwerk-2021", &github.ListMembersOptions{})
	return users
}

func getKeys(users []*github.User) []ssh.PublicKey {
	authKeys := []ssh.PublicKey{}
	for _, user := range users {
		keys, _, _ := client.Users.ListKeys(ctx, user.GetLogin(), &github.ListOptions{})
		for _, key := range keys {
			rsaPubKey, _, _, _, _ := ssh.ParseAuthorizedKey([]byte(key.GetKey()))
			authKeys = append(authKeys, rsaPubKey)
		}
	}
	return authKeys
}
