package main

import (
	"context"
	"io"
	"log"

	"github.com/aalekh12/go-grpc-api/proto"
)

func callSayhelloserverstream(client proto.GreetServiceClient, req *proto.NameList) {
	log.Println("Streaming Started")
	stream, err := client.SayHellloServerStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("error in sending name: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}

		log.Println(msg)
	}

	log.Println("Streaming Finished")
}
