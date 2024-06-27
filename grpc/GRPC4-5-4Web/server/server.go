package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"strings"

	"grpc/hello"

	"log"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Channel(stream hello.HelloService_ChannelServer) error {
	//TODO implement me
	//panic("implement me")
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &hello.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.String) (*hello.String, error) {
	reply := &hello.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

/*
	func main() {
		mux := http.NewServeMux()
		没有启动TLS协议的服务
		h2Handler := h2c.NewHandler(mux, &http2.Server{})
		server := &http.Server{Addr: ":3999", Handler: h2Handler}
		server.ListenAndServe()

		//启动普通的https服务器
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, "hello")
		})

		http.ListenAndServeTLS("localhost:1234", "server.crt", "server.key",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mux.ServeHTTP(w, r)
				return
			}),
		)


		grpcServer := grpc.NewServer()
		hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

		lis, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal(err)
		}
		grpcServer.Serve(lis)
	}
*/
const port = "localhost:8888"

func main() {
	mux := http.NewServeMux()
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

	http.ListenAndServeTLS(port, "server.crt", "server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor != 2 {
				mux.ServeHTTP(w, r)
				return
			}
			if strings.Contains(
				r.Header.Get("Content-Type"), "application/grpc",
			) {
				grpcServer.ServeHTTP(w, r) // gRPC Server
				return
			}

			mux.ServeHTTP(w, r)
			return
		}),
	)
}
