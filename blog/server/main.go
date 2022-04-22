package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ekkinox/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection

func main() {
	fmt.Print("Starting blog server on :50053 ... \n")

	// Mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Failed to create mongo client: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to mongo: %v", err)
	}

	collection = client.Database("blog").Collection("blog")

	// gRPC server
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
