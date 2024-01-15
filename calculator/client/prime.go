package main

import (
	"context"
	"io"
	"log"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

func doGetPrimes(c pb.CalculatorServiceClient) {
	log.Println("doGetPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 210,
	}

	stream, err := c.GetPrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GetPrimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("GetPrimes %s\n", msg.String())
	}
}
