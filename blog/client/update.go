package main

import (
	"context"
	pb "github.com/ekkinox/go-grpc/blog/proto"
	"log"
)

func update(c pb.BlogServiceClient, blog *pb.Blog) {
	log.Println("--Invoking update--")

	blog.AuthorId = "[Updated] " + blog.AuthorId

	_, err := c.Update(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while calling upadte: %v\n", err)
	}

	log.Printf("Update: blog was updated with data: %v\n", blog)
}
