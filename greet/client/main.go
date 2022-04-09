package main

import (
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
	fmt.Printf("Created gRPC client: %v", c)
}
