package watch

import (
	"net/rpc"
	"sync"
)

const KVStoreServiceName = "KVStoreService"

type KVStoreService struct {
	M      map[string]string
	Filter map[string]func(key string)
	mu     sync.Mutex
}

type KVStoreServiceInterface interface {
	NewKVStoreService() *KVStoreService
	Get(key string, value *string) error
	Set(kv [2]string, reply *struct{}) error
	Watch(timeoutSecond int, keyChanged *string) error
}

func RegisterKVStoreService(svc KVStoreServiceInterface) error {
	return rpc.RegisterName(KVStoreServiceName, svc)
}
