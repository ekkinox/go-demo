package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	pb "github.com/ekkinox/go-grpc/blog/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) List(in *emptypb.Empty, stream pb.BlogService_ListServer) error {
	log.Println("--List was invoked--")

	cursor, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, "Cannot list results")
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		data := &BlogItem{}

		if err := cursor.Decode(data); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Cannot decode: %v", err))
		}

		err = stream.Send(convert(data))
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Cannot send: %v", err))
		}
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error during list: %v", err))
	}

	return nil
}
