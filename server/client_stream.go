package main

import (
	"io"
	"log"

	pb "github.com/sajagsubedi/golang-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	log.Println("Streaming started")
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Names %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
