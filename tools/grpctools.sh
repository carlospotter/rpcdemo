#!/bin/bash

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
brew install protobuf

# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative infos.proto
# go get google.golang.org/grpc
# go get google.golang.org/protobuf