package logic

import "github.com/ldaysjun/rpc_learn/protobuf/helloworld"

type greeter struct {
}

func NewGreeter() (helloworld.GreeterServer, error) {
	imp := &greeter{}
	return imp, nil
}
