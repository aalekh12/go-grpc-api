package main

import (
	"context"
	"log"
	"time"

	"github.com/aalekh12/go-grpc-api/proto"
)

func callSayhelloClientStream(client proto.GreetServiceClient, names *proto.NameList) {

	log.Println("Client Streaming Started")
	stream, err := client.SayHellloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("error in streaming %v", err)
	}

	for _, name := range names.Name {
		req := &proto.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}

		log.Printf("Send the request with names %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Println("Streaming Finished")
	if err != nil {
		log.Fatalf("Error While Recving %v", err)
	}
	log.Printf("Recivied Response %v", res)

}
