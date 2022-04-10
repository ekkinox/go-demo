package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) PrimeNumberDecomposition(req *calculatorPb.PrimeNumberDecompositionRequest, stream calculatorPb.CalculatorService_PrimeNumberDecompositionServer) error {

	number := req.GetNumber()

	var k int32
	k = 2

	for number > 1 {
		if number%k == 0 {
			stream.Send(&calculatorPb.PrimeNumberDecompositionResponse{
				Result: k,
			})
			number = number / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func (*server) ComputeAverage(stream calculatorPb.CalculatorService_ComputeAverageServer) error {

	var sum, iter int32
	sum = 0
	iter = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			fmt.Printf("sum: %v, iter: %v", sum, iter)
			return stream.SendAndClose(
				&calculatorPb.ComputeAverageResponse{
					Result: float32(sum) / float32(iter),
				})
		}

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		sum += req.GetNumber()
		iter++
	}
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
