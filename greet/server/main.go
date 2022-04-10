package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	greetPb "github.com/ekkinox/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

type server struct {
	greetPb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetPb.GreetRequest) (*greetPb.GreetResponse, error) {

	title := req.GetGreeting().GetTitle()
	name := req.GetGreeting().GetName()

	fmt.Printf("Greet: received title %s and name %s\n", title, name)

	return &greetPb.GreetResponse{
		Result: fmt.Sprintf("Greetings %s %s!", title, name),
	}, nil
}

func (*server) GreetManyTimes(req *greetPb.GreetManyTimesRequest, stream greetPb.GreetService_GreetManyTimesServer) error {

	title := req.GetGreeting().GetTitle()
	name := req.GetGreeting().GetName()

	fmt.Printf("GreetManyTimes: received title %s and name %s\n", title, name)

	for i := 0; i < 10; i++ {

		resp := &greetPb.GreetManyTimesResponse{
			Result: fmt.Sprintf("Greetings %s %s (#%d)", title, name, i),
		}

		stream.Send(resp)

		time.Sleep(1 * time.Second)
	}

	return nil
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
