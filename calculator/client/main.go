package main

import (
	"context"
	"fmt"
	"io"
	"log"

	calculatorPb "github.com/ekkinox/go-grpc/calculator/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting gRPC client on :50052 ...")

	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorPb.NewCalculatorServiceClient(conn)

	//doSum(c)
	doPrimeNumberDecomposition(c)
}

func doSum(c calculatorPb.CalculatorServiceClient) {
	req := &calculatorPb.Request{
		Integer1: 2,
		Integer2: 3,
	}

	sum, _ := c.Sum(context.Background(), req)

	fmt.Printf("Calculator result: %v", sum.Result)
}

func doPrimeNumberDecomposition(c calculatorPb.CalculatorServiceClient) {

	req := &calculatorPb.PrimeNumberDecompositionRequest{
		Number: 120,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error during server streaming PrimeNumberDecomposition rpc call: %v", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			//end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error during server streaming PrimeNumberDecomposition rpc call: %v", err)
		}

		fmt.Printf("Client server streaming PrimeNumberDecomposition rpc response: %s\n", resp.Result)
	}

	fmt.Printf("Client server streaming PrimeNumberDecomposition finished\n")
}
