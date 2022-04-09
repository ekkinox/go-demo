package main

import (
	"context"
	"fmt"
	"log"

	greetPb "github.com/ekkinox/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting gRPC client on :50051 ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetPb.NewGreetServiceClient(conn)

	doUnary(c)
}

func doUnary(c greetPb.GreetServiceClient) {

	fmt.Println("Starting client unary Greet rpc request ...")

	req := &greetPb.GreetRequest{
		Greeting: &greetPb.Greeting{
			Title: "Mrs",
			Name:  "Jones",
		},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error during unary Greet rpc call: %v", err)
	}

	fmt.Printf("Client unary Greet rpc response: %s", resp.Result)
}
