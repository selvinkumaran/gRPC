package main

import (
	"context"
	"log"

	pb "github.com/selvinkumaran/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {

	log.Println("CreateBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "selvin",
		Title:    "My Forst Blog",
		Content:  "Content of first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("UnExpected Error %v\n", err)
	}

	log.Printf("Blog has been created %s\n", res.Id)

	return res.Id
}
