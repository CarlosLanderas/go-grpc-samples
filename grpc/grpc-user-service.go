package grpc

import (
	"fmt"
	"go-grpc-samples/core"
	"go-grpc-samples/dbclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
)

func Start(address string, db dbclient.BoltClient, credentials credentials.TransportCredentials) {

	userService := core.NewUserService(db)

	var grpcServer *grpc.Server

	if credentials != nil {
		grpcServer = grpc.NewServer(grpc.Creds(credentials))
	} else {
		grpcServer = grpc.NewServer()
	}

	RegisterUserServiceServer(grpcServer, NewUserServiceGrpcServer(userService))

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Error, GRPC service can't listen on port", address)
		os.Exit(1)
	}

	fmt.Println("Starting GRPC service on address", address)
	grpcServer.Serve(lis)
}