package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-grpc-samples/core"
	"go-grpc-samples/dbclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

var(
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func Start(address string, db dbclient.BoltClient, useTLSandAuth bool) {

	userService := core.NewUserService(db)

	var grpcServer *grpc.Server

	if useTLSandAuth {

		cert, err := tls.LoadX509KeyPair("certs/server_cert.pem", "certs/server_key.pem")
		if err != nil {
			log.Fatal(err)
		}

		opts := []grpc.ServerOption {
			grpc.UnaryInterceptor(ensureValidToken),
			grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		}
		grpcServer = grpc.NewServer(opts...)
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

func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	auth := md["authorization"]

	if len(auth) < 1 || auth[0] != "Bearer landetoken" {
		return nil, errInvalidToken
	}

	return handler(ctx, req)
}