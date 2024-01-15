package main

import (
	"context"
	"log"
	"time"

	pb "github.com/piyushdubey/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Piyush"},
		{FirstName: "Zuck"},
		{FirstName: "Elon"},
		{FirstName: "Bill"},
		{FirstName: "Sundar"},
		{FirstName: "Satya"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Could not Long Greet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending Long Greeting req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Received error while long greeting %v\n", err)
	}

	log.Printf("Long Greeting Response = %v\n", res.Result)

}
