package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greetPb "github.com/ekkinox/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

type server struct {
	greetPb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetPb.GreetRequest) (*greetPb.GreetResponse, error) {

	title := req.GetGreeting().GetTitle()
	name := req.GetGreeting().GetName()

	fmt.Printf("received title %s and name %s\n", title, name)

	return &greetPb.GreetResponse{
		Result: fmt.Sprintf("Greetings %s %s!", title, name),
	}, nil
}

func main() {
	fmt.Println("Starting gRPC server on :50051 ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetPb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
