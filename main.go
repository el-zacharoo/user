package main

import (
	"log"
	"net"

	"github.com/el-zacharoo/user/handler"
	pb "github.com/el-zacharoo/user/user.v1"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &handler.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
