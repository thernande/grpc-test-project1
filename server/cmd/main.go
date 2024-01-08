package main

import (
	"log"
	"net"
	"os"

	v1 "github.com/thernande/grpc-test-project1/proto/todo/v1"
	"github.com/thernande/grpc-test-project1/server/internal/controller"
	"github.com/thernande/grpc-test-project1/server/internal/repository/memory"
	"google.golang.org/grpc"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}
	addr := args[0]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(lis)
	log.Printf("listening at %s\n", addr)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	mem := memory.New()
	ctrl := controller.New(mem)
	//registration of endpoints
	v1.RegisterTodoServiceServer(s, ctrl)

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
