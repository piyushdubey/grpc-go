package main

import (
	"io"
	"log"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

func (s *Server) GetAvg(stream pb.CalculatorService_GetAvgServer) error {
	log.Printf("GetAvg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error sending average response: %v\n", err)
		}

		log.Printf("Receiving numbers %d\n", req.Number)
		sum += req.Number
		count++
	}
}
