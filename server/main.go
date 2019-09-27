package main

import (
	"context"
	pb "golaunch/pbmagic"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GiveFileList(ctx context.Context, in *pb.FileListRequest) (*pb.FileListReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.FileListReply{Message: "FileList " + in.GetName()}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGiverServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
