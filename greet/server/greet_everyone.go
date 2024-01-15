package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/piyushdubey/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream. Error: %v\n", err)
		}

		log.Printf("Receiving request: %v\n", req)

		res := fmt.Sprintf("Hello, %v!\n", req.FirstName)

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}
