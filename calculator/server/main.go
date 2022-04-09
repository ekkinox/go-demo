package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculatorPb "github.com/ekkinox/go-grpc/calculator/proto"
	"google.golang.org/grpc"
)

type server struct {
	calculatorPb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorPb.Request) (*calculatorPb.Response, error) {

	log.Printf("Operation: SUM, %v", req)

	return &calculatorPb.Response{
		Result: req.GetInteger1() + req.GetInteger2(),
	}, nil
}

func main() {
	fmt.Println("Starting gRPC server on :50052 ...")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorPb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
