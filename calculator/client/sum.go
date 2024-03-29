package main

import (
	"context"
	"log"

	pb "github.com/selvinkumaran/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("dosun was invoked")

	req := &pb.SumRequest{
		FirstNumber:  12,
		SecondNumber: 1,
	}
	
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("could not sum:%v", err)
	}

	log.Printf("sum %d\n", res.Result)
}
