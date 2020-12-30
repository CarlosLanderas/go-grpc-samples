package main

import (
	"context"
	"go-grpc-samples/dbclient"
	userservicegrpc "go-grpc-samples/grpc"
	"google.golang.org/grpc"
	"testing"
)

func init() {
	db := dbclient.GetDatabase()
	db.OpenDb()
	go userservicegrpc.Start(":8000", db)
}


func BenchmarkGrpcService(b *testing.B) {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		b.Fatal("grpc connection failed", err)
	}
	client := userservicegrpc.NewUserServiceClient(conn)

	for n := 0; n < b.N; n++ {
		ids := []int64{8,6,70}
		client.GetUsers(context.Background(), &userservicegrpc.GetUsersRequest{Ids: ids })
	}
}
