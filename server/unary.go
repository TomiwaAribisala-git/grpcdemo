package main

import (
	"context"

	pb "github.com/TomiwaAribisala-git/grpcdemo/proto"
)

func (s *helloServer) Sayhello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
