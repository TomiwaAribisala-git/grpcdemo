pb:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto

server:
	go run server/*.go

client:
	go run client/*.go