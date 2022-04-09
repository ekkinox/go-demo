# gRPC basics

> Introduction to [gRPC](https://grpc.io) basics with [Go](https://go.dev).

## Requirements

First, you need to follow [gRPC Go quickstart](https://grpc.io/docs/languages/go/quickstart/):

- first install the [protocol buffer compiler (protoc)](https://grpc.io/docs/protoc-installation/) on your system


- then install the protoc Go plugins

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- finally, update your PATH so that the protoc compiler can find the plugins

```shell
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Usage

- to run the Go gRPC server

```shell
$ go run greet/server/main.go
```

- to run the Go gRPC client

```shell
$ go run greet/client/main.go
```

- to (re)generate proto Go stubs

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet/proto/greet.proto
```