package main

import (
	"context"
	"log"
	"time"

	pb "github.com/selvinkumaran/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "selvin"},
		{FirstName: "narkunan"},
		{FirstName: "venkatesan"},
		{FirstName: "kishore"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {

		log.Printf("Sending req %v\n", req)

		stream.Send(req)

		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response fron greet %v\n", err)
	}

	log.Printf("LongGreet %s", res.Result)

}