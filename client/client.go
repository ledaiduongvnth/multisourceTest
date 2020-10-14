package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "multisourceTest/proto"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewIVAClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Signal(ctx, &pb.SignalRequest{Url: "123", Action: "remove"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("response: %s", r.Status)
}
