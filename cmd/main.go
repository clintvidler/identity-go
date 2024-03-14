package main

import (
	"github.com/clintvidler/identity-go/service/server"
)

func main() {
	server.NewServer().Serve()
}
