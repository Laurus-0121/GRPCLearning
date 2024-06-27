package main

import (
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {
}

func main() {
	rpc.Register(new(HelloService))

	for {
		//不主动监听，主动连接外网服务器
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		//基于每个建立的tcp连接提供rpc服务
		rpc.ServeConn(conn)
		conn.Close()
	}
}
