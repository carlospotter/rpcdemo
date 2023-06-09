package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/carlospotter/rpcdemo/broker/infos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/rpc", rpcFunc)
	mux.HandleFunc("/grpc", grpcFunc)

	ctx, cancelCtx := context.WithCancel(context.Background())
	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := srv.ListenAndServe()
		if err == http.ErrServerClosed {
			fmt.Println("server closed")
		}
		if err != nil {
			fmt.Printf("error listening server: %s", err.Error())
		}
		cancelCtx()
	}()

	<-ctx.Done()
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value("serverAddr"))
	io.WriteString(w, "http server up \n")
}

type RPCPayload struct{
	Name string
	Data string
}

func rpcFunc(w http.ResponseWriter, r *http.Request) {
	client, err := rpc.Dial("tcp", "localhost:5001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rpcPayload := RPCPayload{
		Name: "RPC Info",
		Data: time.Now().String(),
	}

	var rpcResult string
	err = client.Call("RPCServer.LogInfo", rpcPayload, &rpcResult) 
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	io.WriteString(w, rpcResult + "\n")
}

func grpcFunc(w http.ResponseWriter, r *http.Request) {
	connection, err := grpc.Dial("localhost:5002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer connection.Close()

	client := infos.NewInfosClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	grpcResult, err := client.LogInfo(ctx, &infos.InfoRequest{
		Info: &infos.Info{
			Name: "gRPC Info",
			Data: time.Now().String(),
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	io.WriteString(w, grpcResult.Result + "\n")
}

