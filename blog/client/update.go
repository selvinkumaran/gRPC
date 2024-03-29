package main

import (
	"context"
	"log"

	pb "github.com/selvinkumaran/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "from update",
		Title:    "new",
		Content:  "this id foe ipdate content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happed while updating: %v\n", err)
	}

	log.Println("Blog was Updated")
}
