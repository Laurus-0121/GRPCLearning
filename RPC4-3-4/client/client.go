package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn net.Conn
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}

			go func() {
				defer conn.Close()

				p := rpc.NewServer()
				p.Register(&HelloService{conn: conn})
				p.ServeConn(conn)
			}()
		}
	}()

	doClientWork(clientChan)
}
func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
