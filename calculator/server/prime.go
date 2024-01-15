package main

import (
	"log"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

func (s *Server) GetPrimes(in *pb.PrimeRequest, stream pb.CalculatorService_GetPrimesServer) error {

	log.Printf("Prime function was invoked with %v\n", in)
	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
