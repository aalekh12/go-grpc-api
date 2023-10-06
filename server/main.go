package main

import (
	"log"
	"net"

	"github.com/aalekh12/go-grpc-api/proto"
	"google.golang.org/grpc"
)

const (
	PORT = ":8080"
)

type helloserver struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("Failed to start the server ", err)
	}

	grpcserver := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcserver, &helloserver{})
	log.Printf("Server started at %v", lis.Addr())
	if err = grpcserver.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}

}
