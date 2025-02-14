package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sajagsubedi/golang-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client Streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitc := make(chan struct{})

	// Goroutine to receive messages
	go func() {
		defer close(waitc) // Ensure channel closes when done
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server finished sending responses")
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			log.Printf("Received: %v", message)
		}
	}()

	// Sending requests
	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending name: %v", err)
		}
		log.Printf("Sent: %s", name)
		time.Sleep(2 * time.Second)
	}

	// Close the sending stream
	stream.CloseSend()

	// Wait for the receive goroutine to finish
	<-waitc

	log.Println("Bidirectional Streaming finished")
}
