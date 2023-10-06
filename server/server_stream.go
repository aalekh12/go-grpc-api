package main

import (
	"log"
	"time"

	"github.com/aalekh12/go-grpc-api/proto"
)

func (s *helloserver) SayHellloServerStreaming(req *proto.NameList, stream proto.GreetService_SayHellloServerStreamingServer) error {
	log.Printf("got request with names: %v ", req.Name)
	for _, v := range req.Name {
		res := &proto.HelloResponse{
			Message: "Hello" + v,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}
	return nil

}
