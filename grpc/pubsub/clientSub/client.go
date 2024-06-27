package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
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
	stream, err := client.Subscribe(context.Background(), &v1.String{Value: "docker:"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
