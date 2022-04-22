package main

import (
	"context"
	pb "github.com/ekkinox/go-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func list(c pb.BlogServiceClient) {
	log.Println("--Invoking list--")

	stream, err := c.List(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling list: %v\n", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while listing stream: %v", err)
		}

		log.Println("Blog: %v", resp)

	}
}
