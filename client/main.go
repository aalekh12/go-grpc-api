package main

import (
	"log"

	"github.com/aalekh12/go-grpc-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start the client %v", err)
	}

	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	req := &proto.NameList{
		Name: []string{"Aalekh", "Ansh", "Ankit"},
	}

	callsayhello(client)
	callSayhelloserverstream(client, req)
	callSayhelloClientStream(client, req)
	callSayhelloBidiretStream(client, req)

}
