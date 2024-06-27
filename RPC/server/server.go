package main

import (
	"hello"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	//TODO implement me
	//panic("implement me")
	*reply = "hello: " + request
	return nil

}

/*func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}*/

func main() {
	/*helloRpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	helloRpc.ServeConn(conn)*/
	//RegisterHelloService(new(HelloService))
	hello.RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Accept error：", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
