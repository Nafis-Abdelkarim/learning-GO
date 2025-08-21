package server

import (
	pb "/grpc/server/addressbook"
	"context"
	"log"
)

type addressbookServer struct{}

func NewAddressBookServer() *addressbookServer {
	return &addressbookServer{}
}

func (s *addressbookServer) ListPersons(ctx context.Context, req *pb.ListPersonsRequest) (*pb.ListPersonsResponse, error) {
	log.Println("Listing persons...")
	return &pb.ListPersonsResponse{}, nil
}
