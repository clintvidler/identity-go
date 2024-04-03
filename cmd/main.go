package main

import (
	"os"

	"github.com/clintvidler/identity-go/app/data"
	"github.com/clintvidler/identity-go/app/server"
)

func main() {
	ds, err := data.NewStore("pg-dev", os.Getenv("DB_DEV_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	if err = ds.Populate(); err != nil {
		panic(err)
	}

	server.NewServer(ds).Serve()
}
