package main

import (
	"context"
	"io"
	"log"

	pb "github.com/TomiwaAribisala-git/grpcdemo/proto"
)

func callsayhelloserverstream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started")
	stream, err := client.Sayhelloserverstreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
}
