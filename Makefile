all:
	@ env GOOS=linux GOARCH=amd64 go build -o proxytunnel cmd/proxy/main.go