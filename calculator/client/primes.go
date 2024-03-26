package main

import (
	"context"
	"io"
	"log"

	pb "github.com/selvinkumaran/grpc-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrime function was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling prime %v\n", err)

	}

	for {

		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading steam %v\n", err)
		}

		log.Printf("the prime is %d\n", res.Result)
	}
}
