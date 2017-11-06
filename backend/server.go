package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbbook "github.com/terut/grpc-sample/proto/book"
	pbuser "github.com/terut/grpc-sample/proto/user"
)

const (
	port = ":9090"
)

type userService struct{}

func (s *userService) Get(ctx context.Context, in *pbuser.UserRequest) (*pbuser.UserResponse, error) {
	log.Println("requested user_id: ", in.Id)
	return &pbuser.UserResponse{Id: 1, Username: "terut"}, nil
}

type bookService struct{}

func (b *bookService) Get(ctx context.Context, in *pbbook.BookRequest) (*pbbook.BookResponse, error) {
	log.Println("requested book_id: ", in.Id)
	return &pbbook.BookResponse{Id: 1, Author: 1, Title: "REST with gRPC"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbuser.RegisterUserServer(s, &userService{})
	pbbook.RegisterBookServer(s, &bookService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
