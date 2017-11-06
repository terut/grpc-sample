package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gwbook "github.com/terut/grpc-sample/proto/book"
	gwuser "github.com/terut/grpc-sample/proto/user"
)

func main() {
	ctx := context.Background()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gwuser.RegisterUserHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Println("failed to register endpoint")
	}
	err = gwbook.RegisterBookHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Println("failed to register endpoint")
	}
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Println("failed to register endpoint")
	}
}
