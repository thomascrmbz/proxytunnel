package client

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"thomascrmbz.com/proxytunnel"
)

type ProxyClient struct {
	proxyServerIP string
	baseCmd       []string
}

func NewProxyClient(proxyServerIP string, proxyServerPort int) *ProxyClient {
	return &ProxyClient{
		proxyServerIP: proxyServerIP,
		baseCmd:       append(baseCmd, proxyServerIP, "-p", strconv.Itoa(proxyServerPort)),
	}
}

var baseCmd = strings.Fields("-tt -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o LogLevel=QUIET -o HostKeyAlgorithms=+ssh-rsa")

func (pc *ProxyClient) execTunnelCmd(cmd proxytunnel.TunnelCmd, agentID int, args ...string) {
	baseCmd = append(append(pc.baseCmd, string(cmd), strconv.Itoa(agentID)), args...)
	// baseCmd = append(pc.baseCmd, args...) // for testing

	exeCmd := exec.Command("ssh", baseCmd...)
	exeCmd.Stdout = os.Stdout
	exeCmd.Stdin = os.Stdin
	exeCmd.Stderr = os.Stderr

	if err := exeCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			code := exitError.ExitCode()
			pc.printErrorMessage(cmd, code)
			os.Exit(code)
		}
	}
}

func (pc *ProxyClient) printErrorMessage(cmd proxytunnel.TunnelCmd, code int) {
	switch code {
	case int(proxytunnel.COMMAND_NOT_FOUND):
		fmt.Println("Command not found")
	case int(proxytunnel.PTY_FAILED):
		fmt.Println("PTY failed to initialize")
	case int(proxytunnel.AGENT_NOT_FOUND):
		fmt.Println("You are trying to connect to a unknown agent")
	case int(proxytunnel.NOT_ALLOWED):
		fmt.Println("You are not authorized for this action")
	case 255:
		fmt.Println("Connection to proxy server refused")
	}
}
