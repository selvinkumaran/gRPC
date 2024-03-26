package main

import (
	"context"
	"log"

	pb "github.com/selvinkumaran/grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	log.Printf("Sum function was invoked with %v", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
