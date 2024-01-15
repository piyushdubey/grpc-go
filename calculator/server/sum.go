package main

import (
	"context"
	"log"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.Operand1 + in.Operand2,
	}, nil
}
