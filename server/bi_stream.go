package main

import (
	"io"
	"log"

	pb "github.com/TomiwaAribisala-git/grpcdemo/proto"
)

func (s *helloServer) Sayhellobidirectionalstreaming(stream pb.GreetService_SayhellobidirectionalstreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
