package main

import (
	"context"
	"log"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

func doGetAverage(c pb.CalculatorServiceClient) {
	log.Println("doGetAverage was invoked")

	stream, err := c.GetAvg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GetAverage %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending numbers: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while opening the stream %v\n", err)
	}

	log.Printf("Average: %v\n", res)
}
