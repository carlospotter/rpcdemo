package main

import (
	"context"
	"log"
	"net/rpc"
)

const rpcPort = "5001"
const grpcPort = "5002"

func main() {
	ctx := context.Background()

	rpcServer := &RPCServer{}
	err := rpc.Register(rpcServer)
	if err != nil {
		log.Panic("unable to register the RPCServer")
	}
	
	go rpcListen()

	go grpcListen()

	<-ctx.Done()
}

