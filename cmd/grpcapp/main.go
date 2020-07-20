package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/ambi/go-web-app-patterns/adapter/rpcserver"
	"github.com/ambi/go-web-app-patterns/api"
	"github.com/ambi/go-web-app-patterns/infra"
	"google.golang.org/grpc"
)

const (
	defaultPort = "50051"
)

func getPort() string {
	port := os.Getenv("PORT")
	i, err := strconv.Atoi(port)
	if err != nil {
		return defaultPort
	}
	if i < 0 || 65535 < i {
		return defaultPort
	}
	return port
}

func main() {
	db, err := infra.NewMySQLDB()
	if err != nil {
		log.Fatal(err)
	}

	srv := rpcserver.NewServer(db)

	lis, err := net.Listen("tcp", ":"+getPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterAPIServer(grpcServer, srv)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
