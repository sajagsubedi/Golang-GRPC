package main

import (
	"io"
	"log"

	pb "github.com/sajagsubedi/golang-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client finished sending messages")
			return nil
		}

		if err != nil {
			return err
		}
		log.Printf("Got names %v", req.Name)

		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

}
