package main

import (
	"log"

	pb "github.com/selvinkumaran/grpc-go/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect:%v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	createBlog(c)

	// readBlog(c, "6606f836e7e4a0f7e0435af2") //valid
	// readBlog(c, "6606f87ecdf3d7b7b7bc678e") //valid
	// readBlog(c, "notId")                    //invalid

	// updateBlog(c, "6606f87ecdf3d7b7b7bc678e")
	// updateBlog(c, "6606f87ecdf3d7b7b7bc678ssqwdwe")

	// listBlogs(c)

	// deleteBlog(c, "6606f836e7e4a0f7e0435af2")
}
