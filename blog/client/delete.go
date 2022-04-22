package main

import (
	"context"
	"log"

	pb "github.com/ekkinox/go-grpc/blog/proto"
)

func delete(c pb.BlogServiceClient, blog *pb.Blog) {
	log.Println("--Invoking delete--")

	_, err := c.Delete(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while calling delete: %v\n", err)
	}

	log.Printf("Delete: blog with id %v was deleted\n", blog.Id)
}
