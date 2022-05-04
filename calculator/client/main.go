package main

import (
	pb "github.com/Chungbien/udemy-grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	clientConnection, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer clientConnection.Close()

	client := pb.NewCalculatorServiceClient(clientConnection)
	log.Printf("Client is: %v", client)
}
