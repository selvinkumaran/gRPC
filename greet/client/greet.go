package main

import (
	"context"
	"log"

	pb "github.com/selvinkumaran/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "selvin",
	})

	if err != nil {
		log.Fatalf("could not greet:%v\n", err)
	}
	log.Printf("Greeting :%s\n", res.Result)
}
