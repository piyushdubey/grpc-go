package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/piyushdubey/grpc-go/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error connecting %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	doCalculate(c)
	doGetPrimes(c)
	doGetAverage(c)
}
