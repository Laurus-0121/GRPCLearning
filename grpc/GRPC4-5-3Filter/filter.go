package GRPC4_5_3Filter

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func filter(
	//截取器增加了对 gRPC 方法异常的捕获：
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("filter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

/*以下是 go-grpc-middleware 包中链式截取器的简单用法



myServer := grpc.NewServer(
grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
filter1, filter2, ...
)),
grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
filter1, filter2, ...
)),
)*/
