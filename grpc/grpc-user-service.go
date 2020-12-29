package grpc

import (
	"fmt"
	"go-grpc-samples/core"
	"go-grpc-samples/dbclient"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func Start(address string, db dbclient.BoltClient) {

	userService := core.NewUserService(db)

	grpcServer := grpc.NewServer()
	RegisterUserServiceServer(grpcServer, NewUserServiceGrpcServer(userService))

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Error, GRPC service can't listen on port %s", address)
		os.Exit(1)
	}

	fmt.Println("Starting GRPC service on address", address)
	grpcServer.Serve(lis)
}