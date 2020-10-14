package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	iva "multisourceTest/pkg"
	pb "multisourceTest/proto"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedIVAServer
	iva *iva.IVA
}

func (s *server) Signal(ctx context.Context, request *pb.SignalRequest) (*pb.SignalReply, error) {
	log.Printf(request.Url, request.Action)
	if request.Action == "add" {
		s.iva.AddSource(request.Url)
	} else if request.Action == "remove" {
		s.iva.RemoveSource(request.Url)
	}
	return &pb.SignalReply{Status: "0k"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	newIVA := iva.NewIVA()
	pb.RegisterIVAServer(s, &server{iva: newIVA})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
