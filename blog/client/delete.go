package main

import (
	"context"
	"log"

	pb "github.com/selvinkumaran/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {

	log.Println("deleteBlog was invoked")

	req := &pb.BlogId{Id: id}

	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was Deleted succesfully")
}
