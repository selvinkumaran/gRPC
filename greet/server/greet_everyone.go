package main

import (
	"io"
	"log"

	pb "github.com/selvinkumaran/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Println("GreetEveryone was invoked")

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading fron client :%v\n", err)
		}

		res := "Hello " + req.FirstName + " !"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while readong fron client: %v\n", err)
		}
	}
}
