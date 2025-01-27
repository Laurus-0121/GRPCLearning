package hello

import "net/rpc"

const HelloServiceName = "HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
	isLogin(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
