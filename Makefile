all:
	env GOOS=linux go build -o proxy-gateway-new cmd/proxy/main.go && rsync ./proxy-gateway-new devops@192.168.10.13:~