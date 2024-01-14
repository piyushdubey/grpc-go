package main

import (
	"log"
	"net"
)

var addr string = "localhost:50051"

func main() {
	conn, err := net.Dial(addr, "")

	if err != nil {
		log.Fatalf("Error connecting %v\n", err)
	}

	defer conn.Close()
}
