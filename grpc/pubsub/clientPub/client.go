package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	v1 "pubsub/pubsub/common/v1"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &v1.String{Value: "golang:hello Golang!"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &v1.String{Value: "docker:hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
