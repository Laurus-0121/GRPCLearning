package main

import (
	"fmt"
	"hello"
	"log"
	"net/rpc"
)

/*func main() {
	clientSub, err := helloRpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = clientSub.Call("HelloService.Hello", "nihao", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
*/

type HelloServiceClient struct {
	*rpc.Client
}

func (p HelloServiceClient) Hello(request string, reply *string) error {
	//TODO implement me
	//panic("implement me")
	return p.Client.Call(hello.HelloServiceName+".Hello", request, reply)
}

// 类型断言，保证client能够实现interface里面所有的函数
var _ hello.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func main() {
	/*clientSub, err := helloRpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = clientSub.Call(server.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}*/
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(reply)
}
