pb:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto

builds:
	go build -o server/*.go

buildc:
	go build -o client/*.go

runs:
	go run server/*.go

runc:
	go run client/*.go