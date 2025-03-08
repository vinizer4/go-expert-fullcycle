package main

import (
	"github.com/goexpert/challenges/client-server-api/client"
	"github.com/goexpert/challenges/client-server-api/server"
)

func main() {
	go server.RunServer()
	client.RunClient()
}
