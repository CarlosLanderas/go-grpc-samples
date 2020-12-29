package main

import (
	"flag"
	"go-grpc-samples/dbclient"
)


func main() {
	seed := flag.Bool("seed", true, "Seed database accounts")

	flag.Parse()

	db := dbclient.NewDatabase()

	if *seed {
		db.Seed()
	}

	db.OpenDb()

}
