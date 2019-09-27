package main

import (
	"context"
	"flag"
	pb "golaunch/pbmagic"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) GiveFileList(ctx context.Context, in *pb.FileListRequest) (*pb.FileListReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.FileListReply{Message: "FileList " + in.GetName()}, nil
}

func main() {
	address	:= flag.String("port", ":50051", "Default address")
	flag.Parse()
	listener, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGiverServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
