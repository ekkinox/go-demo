package main

import (
	"fmt"
	"log"

	pb "github.com/ekkinox/go-grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Using blog client on :50053 ...")

	conn, err := grpc.Dial(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	fmt.Println("done !")

	bid := create(c)
	b := read(c, bid)
	update(c, b)

	list(c)

	delete(c, b)
}
