package main

import (
	userservicegrpc "go-grpc-samples/grpc"
)


func main() {

	userservicegrpc.Start(":8000")
}
