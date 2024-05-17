package main

import (
	"io"
	"log"

	pb "github.com/TomiwaAribisala-git/grpcdemo/proto"
)

func (s *helloServer) Sayhelloclientstreaming(stream pb.GreetService_SayhelloclientstreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
