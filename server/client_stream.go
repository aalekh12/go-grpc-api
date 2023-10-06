package main

import (
	"io"
	"log"

	"github.com/aalekh12/go-grpc-api/proto"
)

func (s *helloserver) SayHellloClientStreaming(stream proto.GreetService_SayHellloClientStreamingServer) error {

	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.MessageList{Message: messages})
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}

		log.Printf("Got request with names %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
