package main

import (
	"fmt"
	"go-grpc-samples/dbclient"
	userservicegrpc "go-grpc-samples/grpc"
	"go-grpc-samples/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	db := dbclient.GetDatabase()
	db.OpenDb()

	go func() {
		userservicegrpc.Start(":8000", db)
	}()

	go func() {
		server := http.NewServer(db)
		server.Start(":8001")
	}()


	<-c
	fmt.Println("Server shutdown...")
}
