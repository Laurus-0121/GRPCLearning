package main

import (
	"fmt"
	"hello"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func (p *HelloService) Login(request string, reply *string) error {
	//*reply = "hello:" + request + ",from" + p.conn.RemoteAddr().String()
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	//*reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ",from" + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	err := hello.RegisterHelloService(new(HelloService))

	if err != nil {
		return
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Accept error: ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		go rpc.ServeConn(conn)
	}
}
