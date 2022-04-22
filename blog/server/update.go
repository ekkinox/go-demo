package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	pb "github.com/ekkinox/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Update(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("--Update was invoked: %v--\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse id: %v", err))
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot update for id %v: %v", oid, err))
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find for id %v: %v", oid, err))
	}

	return &emptypb.Empty{}, nil
}
