package main

import (
	"context"
	"io"
	"log"

	pb "github.com/selvinkumaran/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "selvin",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling from GreetManyTimes %v\n", err)
	}

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}

		log.Printf("GreetManyTimes : %s\n", msg.Result)
	}
}