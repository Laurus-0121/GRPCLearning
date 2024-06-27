package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
	"watch"
)

type KVStoreServiceClient struct {
	*rpc.Client
}

func (K KVStoreServiceClient) NewKVStoreService() *watch.KVStoreService {
	//TODO implement me
	//panic("implement me")
	return K.NewKVStoreService()
}

func (K KVStoreServiceClient) Get(key string, value *string) error {
	//TODO implement me
	//panic("implement me")
	return K.Get(key, value)
}

func (K KVStoreServiceClient) Set(kv [2]string, reply *struct{}) error {
	//TODO implement me
	//panic("implement me")
	return K.Set(kv, reply)
}

func (K KVStoreServiceClient) Watch(timeoutSecond int, keyChanged *string) error {
	//TODO implement me
	//panic("implement me")
	return K.Watch(timeoutSecond, keyChanged)
}

var _ watch.KVStoreServiceInterface = (*KVStoreServiceClient)(nil)

func DialKVStoreService(network, address string) (*KVStoreServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("DialKVStoreService fail: ", err)
	}
	return &KVStoreServiceClient{Client: c}, nil
}

func main() {
	client, err := DialKVStoreService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	doClientWork(client.Client)

}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch: ", &keyChanged)
	}()

	err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("ok!!")

}
