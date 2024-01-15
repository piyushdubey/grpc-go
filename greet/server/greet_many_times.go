package main

import (
	"fmt"
	"log"

	pb "github.com/piyushdubey/grpc-go/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, Number: %d\n", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
