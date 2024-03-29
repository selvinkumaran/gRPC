package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/selvinkumaran/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	
	log.Println("doMax Was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{2, 5, 1, 8, 3, 13}

		for _, number := range numbers {
			log.Printf("Sending number:%v\n", number)

			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving stream:%v\n", err)
				break
			}
			log.Printf("Received :%v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
