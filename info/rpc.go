package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

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

func rpcListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		log.Println(err)
	}

	defer listen.Close()

	log.Printf("RPC Server started on port %s", rpcPort)

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}