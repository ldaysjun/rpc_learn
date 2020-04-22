package logic

import (
	"fmt"
	"github.com/ldaysjun/rpc_learn/protobuf/helloworld"
	"golang.org/x/net/context"
)

func (g *greeter) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	rsp := &helloworld.HelloReply{
		Message: fmt.Sprintf("hello:%s", req.Name),
	}
	return rsp, nil
}
