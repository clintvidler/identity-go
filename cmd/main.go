package main

import (
	"os"

	"github.com/clintvidler/identity-go/app/data"
	"github.com/clintvidler/identity-go/app/server"
	"github.com/clintvidler/identity-go/services"
)

func main() {
	ds, err := data.NewStore("pg-dev", os.Getenv("DB_DEV_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	if err = ds.Populate(); err != nil {
		panic(err)
	}

	ec := services.NewEmailClient("host.docker.internal", "7399")

	server.NewServer(ds, ec).Serve()
}
