package main

import (
	"flag"
	"fmt"
	"go-grpc-samples/dbclient"
	userservicegrpc "go-grpc-samples/grpc"
	"go-grpc-samples/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	seed := flag.Bool("seed", false, "Enable database seed")
	flag.Parse()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	db := dbclient.GetDatabase()
	db.OpenDb()

	if *seed {
		db.Seed()
	}

	go userservicegrpc.Start(":8000", db, true)

	go func() {

		server := http.NewServer(db)
		server.Start(":8001")
	}()


	<-c
	fmt.Println("Server shutdown...")
}
