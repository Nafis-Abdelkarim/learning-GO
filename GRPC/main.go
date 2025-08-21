package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Cloud not listen on port")
	}

	grpcServer := grpc.NewServer()
}
