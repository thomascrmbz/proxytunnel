package handler

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gliderlabs/ssh"
	"github.com/phayes/freeport"
	"thomascrmbz.com/proxytunnel/agent"
)

var api *cloudflare.API

func init() {
	api, _ = cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
}

func ProxyHandler(a *agent.Agent, s ssh.Session) {
	port := s.Command()[2]
	proxyport := getFreePort()
	conn := a.IP + " -p " + strconv.Itoa(a.Port)
	recordID := ""
	sshExe(a, s, sshOptions{
		CopyStdin:  true,
		CopyStdout: false,
		Cmd:        strings.Fields("-N -L pulu.trikthom.com:" + proxyport + ":localhost:" + port + " " + conn),
		OnRun: func() {
			s.Write([]byte("Proxied " + a.IP + ":" + port + " to pulu.trikthom.com:" + proxyport + "\n"))
			go func() {
				recordID = setupDNS(s, proxyport)
			}()
		},
		OnDone: func() {
			s.Write([]byte("Stopped proxy\n"))
			if len(recordID) > 0 {
				go removeDNS(recordID)
			}
		},
	})

}

func getFreePort() string {
	port, _ := freeport.GetFreePort()
	return strconv.Itoa(port)
}

func setupDNS(s ssh.Session, proxyport string) string {
	time.Sleep(time.Second)
	_, err := exec.Command("curl", strings.Fields("-I pulu.trikthom.com:"+proxyport)...).Output()
	if err == nil {
		s.Write([]byte("Proxy is a HTTP server\n"))
		s.Write([]byte("Setting up HTTPS..."))

		res, err := api.CreateDNSRecord(context.Background(), "7fe96359f5ed714813f8130de121256d", cloudflare.DNSRecord{
			Type:    "CNAME",
			Name:    "proxy-" + proxyport,
			Content: "pulu.trikthom.com",
			Proxied: &[]bool{true}[0],
		})
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			s.Write([]byte("failed\n"))
		} else {
			s.Write([]byte("done\n"))
			s.Write([]byte("Visit: https://proxy-" + proxyport + ".trikthom.com/\n"))
		}

		return res.Result.ID
	}
	return ""
}

func removeDNS(recordID string) {
	err := api.DeleteDNSRecord(context.Background(), "7fe96359f5ed714813f8130de121256d", recordID)
	if err != nil {
		fmt.Println(err)
	}
}
