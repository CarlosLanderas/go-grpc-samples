package main

import (
	"flag"
	"go-grpc-samples/dbclient"
	userservicegrpc "go-grpc-samples/grpc"
	user_service "go-grpc-samples/user-service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)


func main() {
	seed := flag.Bool("seed", true, "Seed database accounts")

	flag.Parse()

	db := dbclient.NewDatabase()

	if *seed {
		db.Seed()
	}

	db.OpenDb()
	userService := user_service.NewUserService(db)

	grpcServer := grpc.NewServer()
	userservicegrpc.RegisterUserServiceServer(grpcServer, userservicegrpc.NewUserServiceGrpcServer(userService))

	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal("Error, GRPC service can't listen on port 8000")
		os.Exit(1)
	}

	grpcServer.Serve(lis)
}
