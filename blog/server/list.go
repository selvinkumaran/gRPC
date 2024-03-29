package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/selvinkumaran/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	
	log.Println("ListBlogs was invoked")

	curr, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal Error: %v\n", err),
		)
	}

	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		data := &BlogItem{}
		err := curr.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding error from mongodb: %v\n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal error: %v\n", err),
		)
	}
	return nil
}
