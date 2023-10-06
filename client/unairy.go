package main

import (
	"context"
	"log"
	"time"

	"github.com/aalekh12/go-grpc-api/proto"
)

func callsayhello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("Error in Respnse %v", err)
	}

	log.Printf("%s", resp.Message)

}
