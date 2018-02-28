package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Pepeye/microed/consignments/service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// Repository interface
type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	List() []*pb.Consignment
}

// Repository type
type Repository struct {
	consignments []*pb.Consignment
}

// Create method creates a new consingment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// List method returns a slice of consignments
func (repo *Repository) List() []*pb.Consignment {
	return repo.consignments
}

// service struct type
type service struct {
	repo repository
}

// CreateConsignment - method in protobuf definition
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// save the consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// return matching the `Response` message in protobuf definition
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignment(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.List()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen:")
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
