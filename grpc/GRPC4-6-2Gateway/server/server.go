package main

import (
	"context"
	v1 "gateway/Gateway/common/gateway/v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"runtime"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	server := grpc.NewServer()
	srv := v1.UnimplementedRestServiceServer{}
	err := v1.RegisterRestServiceServer()

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", mux)
}
