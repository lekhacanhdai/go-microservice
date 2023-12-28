package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcserver/generated"
	"log"
	"net"
)

type server1 struct {
	generated.UnimplementedHelloServiceServer
}

func (s *server1) HelloWorld(ctx context.Context, req *generated.HelloRequest) (*generated.HelloResponse, error) {
	return &generated.HelloResponse{
		HelloWorld: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	generated.RegisterHelloServiceServer(s, &server1{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
