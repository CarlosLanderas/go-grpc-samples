package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"go-grpc-samples/dbclient"
	userservicegrpc "go-grpc-samples/grpc"
	"go-grpc-samples/http"
	"google.golang.org/grpc/credentials"
	"log"
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


	cert, err := tls.LoadX509KeyPair("certs/server_cert.pem", "certs/server_key.pem")
	if err != nil {
		log.Fatal(err)
	}


	go userservicegrpc.Start(":8000", db, credentials.NewServerTLSFromCert(&cert))

	go func() {

		server := http.NewServer(db)
		server.Start(":8001")
	}()


	<-c
	fmt.Println("Server shutdown...")
}
