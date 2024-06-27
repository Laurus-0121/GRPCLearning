package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	v1 "grpc/hello"
	"io/ioutil"
	"log"
)

func main() {
	//客户端证书验证
	//creds, err := credentials.NewClientTLSFromFile("server.crt", "server.grpc.io")
	certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}
	cretPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := cretPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "server.grpc.io",
		RootCAs:      cretPool,
	})

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &v1.String{Value: "hi"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	//stream,err := client.Channel(context.Background())

}
