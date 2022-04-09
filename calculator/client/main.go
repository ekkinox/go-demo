package main

import (
	"context"
	"fmt"
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

	doSum(c)
}

func doSum(c calculatorPb.CalculatorServiceClient) {
	req := &calculatorPb.Request{
		Integer1: 2,
		Integer2: 3,
	}

	sum, _ := c.Sum(context.Background(), req)

	fmt.Printf("Calculator result: %v", sum.Result)
}
