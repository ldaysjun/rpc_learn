package main

import (
	"github.com/ldaysjun/rpc_learn/protobuf/helloworld"
	"github.com/ldaysjun/rpc_learn/rpc_demo/internal/logic"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50052"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	imp, err := logic.NewGreeter()
	if err != nil {
		panic(err)
	}
	helloworld.RegisterGreeterServer(s, imp)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
