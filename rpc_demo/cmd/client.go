package main

import (
	"context"
	"fmt"
	"github.com/ldaysjun/rpc_learn/protobuf/helloworld"
	"github.com/ldaysjun/rpc_learn/rpc_demo/pkg/api"
	"log"
	"time"
)

func main() {
	client := api.NewGreeterClient()
	req := &helloworld.HelloRequest{
		Name: "ldaysjun",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rsp, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}



	fmt.Println("rsp.message = ", rsp.Message)
}
