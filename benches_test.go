package main

import (
	"bytes"
	"context"
	"encoding/json"
	"go-grpc-samples/dbclient"
	userservicegrpc "go-grpc-samples/grpc"
	"go-grpc-samples/http"
	"google.golang.org/grpc"
	net "net/http"
	"testing"
)

var db dbclient.BoltClient

func init() {
	db = dbclient.GetDatabase()
	db.OpenDb()
	go userservicegrpc.Start(":8000", db)
	go func() {
		server := http.NewServer(db)
		server.Start(":8001")
	}()
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

func BenchmarkHttpService(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ids := []int64{8,6,70}
		content, _ := json.Marshal(ids)
		reader := bytes.NewReader(content)

		net.DefaultClient.Post("http://localhost:8001/users", "application/json", reader)
	}
}
