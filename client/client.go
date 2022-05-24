package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nXnUs25/gRPCSum/proto"
)

var addr string = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewSumServiceClient(conn)
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 6,
		B: 5,
	})
	if err != nil {
		log.Fatalf("Failed to greeting: %v", err)
	}

	log.Printf("%v", res.Result)
}
