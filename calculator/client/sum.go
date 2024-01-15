package main

import (
	"context"
	"log"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Operand1: 2,
		Operand2: 9,
	})

	if err != nil {
		log.Fatalf("Could not calculate %v\n", err)
	}

	log.Printf("Calculator: %d\n", res.Result)
}
