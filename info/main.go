package main

import (
	"context"
	"fmt"
	"log"
	"net"
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

func rpcListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		log.Println(err)
	}

	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}

// RPC:
type RPCServer struct{}

type RPCPayload struct{
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, response *string) error {
	log.Printf("%s - %s", payload.Name, payload.Data)

	*response = fmt.Sprintf("%s processed log via RPC", payload.Name)
	return nil
}