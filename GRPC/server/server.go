package server

import (
	"context"
	"log"

	pb "grpc/server/addressbook"
)

type addressbookServer struct {
}

func NewAddressbookServer() *addressbookServer {
	return &addressbookServer{}
}

func (s *addressbookServer) ListPersons(context context.Context, req *pb.ListPersonsRequest) (*pb.ListPersonsResponses, error) {
	log.Println("Listing perons...")
	return &pb.ListPersonsResponses{}, nil
}
