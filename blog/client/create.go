package main

import (
	"context"
	pb "github.com/ekkinox/go-grpc/blog/proto"
	"log"
)

func create(c pb.BlogServiceClient) *pb.BlogId {
	log.Println("--Invoking create--")

	blog := &pb.Blog{
		AuthorId: "Jonathan",
		Title:    "Example",
		Content:  "Some example of blog item",
	}

	res, err := c.Create(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while calling create: %v\n", err)
	}

	log.Printf("Create: blog was created with id: %v\n", res.Id)

	return res
}
