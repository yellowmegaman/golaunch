package main

import (
	"context"
	pb "golaunch/pbmagic"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()
	c := pb.NewGiverClient(connection)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GiveFileList(ctx, &pb.FileListRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not get filelist: %v", err)
	}
	log.Printf("Gettin: %s", r.GetMessage())
}
