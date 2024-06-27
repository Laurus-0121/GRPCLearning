package main

import (
	"context"
	"github.com/moby/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	v1 "pubsub/pubsub/common/v1"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{pub: pubsub.NewPublisher(100*time.Millisecond, 10)}
}

func (p *PubsubService) Publish(ctx context.Context, arg *v1.String) (*v1.String, error) {
	p.pub.Publish(arg.GetValue())
	return &v1.String{}, nil
}
func (p *PubsubService) Subscribe(args *v1.String, stream v1.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, args.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&v1.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	server := grpc.NewServer()
	v1.RegisterPubsubServiceServer(server, NewPubsubService())

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve(listen)
	if err != nil {
		return
	}
}
