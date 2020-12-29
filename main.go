package main

import (
	userservicegrpc "go-grpc-samples/grpc"
	"go-grpc-samples/http"
)


func main() {

	go func() {
		server := http.NewServer()
		server.Start(":8001")
	}()

	userservicegrpc.Start(":8000")
}
