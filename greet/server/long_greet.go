package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/piyushdubey/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("Reached EOF white Long Greeting. Response: %v", res)
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while LongGreeting. Error: %v\n", err)
		}

		log.Printf("Receiving request: %v\n", req)
		res += fmt.Sprintf("Hello, %v\n", req.FirstName)
	}
}
