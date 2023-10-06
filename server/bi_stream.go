package main

import (
	"io"
	"log"

	"github.com/aalekh12/go-grpc-api/proto"
)

func (s *helloserver) SayHellloBiderctionalStreaming(stream proto.GreetService_SayHellloBiderctionalStreamingServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		log.Printf("got request with names %v", req.Name)
		res := &proto.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

}
