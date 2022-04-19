package main

import (
	pb "github.com/ekkinox/go-grpc/blog/proto/blog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func convert(b *BlogItem) pb.Blog {

}
