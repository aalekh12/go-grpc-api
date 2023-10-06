package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/aalekh12/go-grpc-api/proto"
)

func callSayhelloBidiretStream(client proto.GreetServiceClient, names *proto.NameList) {
	log.Println("Bidirectonal Streaming Started")

	stream, err := client.SayHellloBiderctionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("error in stremaing", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Name {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error in sending request %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Biredectional Streaming Finished")
}
