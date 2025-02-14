package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sajagsubedi/golang-grpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive message %v", err)
		}
		log.Println(message)
	}
	log.Println("Streaming finished")
}
