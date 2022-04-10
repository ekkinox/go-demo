package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
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

func doServerStreaming(c greetPb.GreetServiceClient) {

	fmt.Println("Starting client server Streaming GreetManyTimes rpc request ...")

	req := &greetPb.GreetManyTimesRequest{
		Greeting: &greetPb.Greeting{
			Title: "Mrs",
			Name:  "Jones",
		},
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error during server streaming GreetManyTimes rpc call: %v", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			//end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error during server streaming GreetManyTimes rpc call: %v", err)
		}

		fmt.Printf("Client server streaming GreetManyTimes rpc response: %s\n", resp.Result)
	}

	fmt.Printf("Client server streaming GreetManyTimes finished\n")
}

func doClientStreaming(c greetPb.GreetServiceClient) {

	fmt.Println("Starting client server Streaming LongGreet rpc request ...")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error during client streaming LongGreet rpc call: %v", err)
	}

	reqs := []*greetPb.LongGreetRequest{
		{
			Greeting: &greetPb.Greeting{
				Title: "Mrs",
				Name:  "Doe",
			},
		},
		{
			Greeting: &greetPb.Greeting{
				Title: "Mr",
				Name:  "Du",
			},
		},
		{
			Greeting: &greetPb.Greeting{
				Title: "Mrs",
				Name:  "Po",
			},
		},
		{
			Greeting: &greetPb.Greeting{
				Title: "Mr",
				Name:  "Sa",
			},
		},
	}

	for _, req := range reqs {
		fmt.Printf("Sending: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error during client CloseAndRecv LongGreet rpc call: %v", err)
	}

	fmt.Printf("Server response: %v", resp.Result)
}

func doBiDiStreaming(c greetPb.GreetServiceClient) {

	fmt.Println("Starting BiDi streaming ...")

	stream, err := c.GreetAll(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	wait := make(chan struct{})

	reqs := []*greetPb.GreetAllRequest{
		{
			Greeting: &greetPb.Greeting{
				Title: "Mrs",
				Name:  "Doe",
			},
		},
		{
			Greeting: &greetPb.Greeting{
				Title: "Mr",
				Name:  "Du",
			},
		},
		{
			Greeting: &greetPb.Greeting{
				Title: "Mrs",
				Name:  "Po",
			},
		},
		{
			Greeting: &greetPb.Greeting{
				Title: "Mr",
				Name:  "Sa",
			},
		},
	}

	//senders
	go func() {
		for _, req := range reqs {
			fmt.Printf("Sending: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)

		}
		stream.CloseSend()
	}()

	//receivers
	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("error: %v", err)
				break
			}

			fmt.Printf("Received: %v\n", resp.Result)
		}
		close(wait)
	}()

	<-wait
}
