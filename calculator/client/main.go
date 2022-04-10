package main

import (
	"context"
	"fmt"
	calculatorPb "github.com/ekkinox/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"io"
	"log"
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
	//doPrimeNumberDecomposition(c)
	//doComputeAverage(c)
	doFindMax(c)
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

func doComputeAverage(c calculatorPb.CalculatorServiceClient) {
	reqs := []*calculatorPb.ComputeAverageRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	resp, err := stream.CloseAndRecv()
	fmt.Printf("Average: %v", resp.Result)
}

func doFindMax(c calculatorPb.CalculatorServiceClient) {

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	reqs := []*calculatorPb.FindMaxRequest{
		{
			Number: 1,
		},
		{
			Number: 5,
		},
		{
			Number: 3,
		},
		{
			Number: 6,
		},
		{
			Number: 2,
		},
		{
			Number: 20,
		},
	}

	wait := make(chan struct{})

	//senders
	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	//reveivers
	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				close(wait)
				log.Fatalf("error: %v", err)
			}

			log.Println(resp.Max)

		}
		close(wait)
	}()

	<-wait
}
