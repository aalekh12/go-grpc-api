package main

import (
	"context"

	"github.com/aalekh12/go-grpc-api/proto"
)

func (s *helloserver) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {

	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
