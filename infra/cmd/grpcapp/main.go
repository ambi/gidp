package main

import (
	"log"
	"net"

	"github.com/ambi/gidp/adapter/rpccontroller"
	"github.com/ambi/gidp/infra"
	"google.golang.org/grpc"
)

const (
	defaultPort = "50051"
)

func main() {
	db, err := infra.NewMySQLDB()
	if err != nil {
		log.Fatal(err)
	}

	srv := rpccontroller.NewServer(db)

	lis, err := net.Listen("tcp", ":"+infra.GetPort(defaultPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	rpccontroller.RegisterAPIServer(grpcSrv, srv)

	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
