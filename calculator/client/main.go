package main

import (
	"context"
	"fmt"
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
	getSum(client)
}

func getSum(c pb.CalculatorServiceClient) {
	fmt.Printf("Get Sum")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		NumberA: 10,
		NumberB: 20,
	})

	if err != nil {
		log.Fatalf("Can not get sum: %v", err)
	}
	log.Printf("Client is: %v", res.Result)
}
