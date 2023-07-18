package main

import (
	server "jemoeder.com/cmd/api"
)

func main() {
	server, err := server.NewServer("127.0.0.1:3333")
	if err != nil {
		panic(err)
	}
	server.Start()
}
