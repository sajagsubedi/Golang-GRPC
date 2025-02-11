package main

import (
	"log"

	pb "github.com/sajagsubedi/golang-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080" // Port to listen on

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	// names := &pb.NamesList{
	// 	Names: []string{"Sajag", "Ram"},
	// }
	callSayHello(client)
}
