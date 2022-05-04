package main

import (
	"context"
	"fmt"
	pb "github.com/Chungbien/udemy-grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", request)
	firstNumber := request.NumberA
	secondNumber := request.NumberB
	sum := firstNumber + secondNumber
	res := &pb.SumResponse{
		Result: int64(sum),
	}
	return res, nil
}

func main() {
	fmt.Printf("Serve on 127.0.0.1:8080")
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Err:", err)
	}
	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Err:", err)
	}
}
