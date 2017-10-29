package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/terut/grpc-sample/proto"
)

func main() {
	ctx := context.Background()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterUserHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Pritnln("failed to register endpoint")
	}
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Pritnln("failed to register endpoint")
	}
}
