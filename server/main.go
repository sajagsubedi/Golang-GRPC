package main

import (
	"log"
	"net"

	pb "github.com/sajagsubedi/golang-grpc/proto"
	"google.golang.org/grpc"
)

const port = ":8080"

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	grpcServer := grpc.NewServer()

	//register the service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server started on  %s", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving %v", err)
	}

}
