### Golang gRPC demo
- Unary API
- Server Streaming
- Client Streaming
- Bi-directional Streaming

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

```sh
apt install protobuf-compiler
```

```sh
protoc --version
```

```sh
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

```sh   // client
go run *.go
```

```sh   // server
go run *.go
```