package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
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

func rpcFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /rpc request\n", ctx.Value("serverAddr"))
	io.WriteString(w, "http server up: route rpc\n")
}

func grpcFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /grpc request\n", ctx.Value("serverAddr"))
	io.WriteString(w, "http server up: route grpc\n")
}
