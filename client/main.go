package main

import (
	"log"

	pb "github.com/TomiwaAribisala-git/grpcdemo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	//callsayhello(client)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}
	callSayHelloServerStream(client, names)

}
