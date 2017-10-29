package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/terut/grpc-sample/proto"
)

const (
	port = ":9090"
)

type userService struct{}

func (s *userService) Get(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("requested id: ", in.Id)
	return &pb.UserResponse{Id: 1, Username: "terut"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &userService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
