// Package main implements a server for Greeter service.
package main

import (
"context"
"log"
"net"

pb "github.com/shgopher/gRPC-demo/hello"
"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.Apply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Apply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
