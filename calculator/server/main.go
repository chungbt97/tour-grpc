package main

import (
	"fmt"
	calulatorPb "github.com/Chungbien/udemy-grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	calulatorPb.UnimplementedCalculatorServiceServer
}

func main() {
	fmt.Printf("Serve on 127.0.0.1:8080")
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Err:", err)
	}
	s := grpc.NewServer()

	calulatorPb.RegisterCalculatorServiceServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Err:", err)
	}
}
