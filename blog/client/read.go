package main

import (
	"context"
	pb "github.com/ekkinox/go-grpc/blog/proto"
	"log"
)

func read(c pb.BlogServiceClient, id *pb.BlogId) *pb.Blog {
	log.Println("--Invoking read--")

	blogId := &pb.BlogId{
		Id: id.Id,
	}

	res, err := c.Read(context.Background(), blogId)
	if err != nil {
		log.Fatalf("Error while calling read: %v\n", err)
	}

	log.Printf("Read: blog was found with data: %v\n", res)

	return res
}
