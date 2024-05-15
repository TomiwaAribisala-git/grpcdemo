package main

import (
	"context"
	"log"
	"time"

	pb "github.com/TomiwaAribisala-git/grpcdemo/proto"
)

func callsayhello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Sayhello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
