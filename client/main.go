package main

import (
	"context"
	"flag"
	pb "golaunch/pbmagic"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	address := flag.String("address", "127.0.0.1:50051", "Default host")
	name := flag.String("name", "default", "Default name")
	flag.Parse()
	connection, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()
	c := pb.NewGiverClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GiveFileList(ctx, &pb.FileListRequest{Name: *name})
	if err != nil {
		log.Fatalf("Could not get filelist: %v", err)
	}
	log.Printf("Gettin: %s", r.GetMessage())
}
