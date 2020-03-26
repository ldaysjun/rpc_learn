package api

import (
	"github.com/ldaysjun/rpc_learn/protobuf/helloworld"
	"google.golang.org/grpc"
	"log"
)

var address = "localhost:50051"

func NewGreeterClient() helloworld.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	return c
}
