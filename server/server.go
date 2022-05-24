package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nXnUs25/gRPCSum/proto"
	"google.golang.org/grpc"
)

type SumServer struct {
	pb.SumServiceServer
}

var addrs string = "0.0.0.0:50051"

func (s *SumServer) Sum(cxt context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum func called %v\n", in)

	return &pb.SumResponse{Result: in.A + in.B}, nil
}

func main() {
	listen, err := net.Listen("tcp", addrs)

	if err != nil {
		log.Fatalf("Failed to listen on %s - %v", addrs, err)
	}

	log.Printf("listening on %s\n", addrs)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &SumServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on %s - %v", addrs, err)
	}
}
