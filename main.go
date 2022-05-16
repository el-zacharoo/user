package main

import (
	"log"
	"net"

	"github.com/el-zacharoo/user/handler"
	"github.com/el-zacharoo/user/store"
	pb "github.com/el-zacharoo/user/user.v1"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	st := store.Connect()
	grpcServer := grpc.NewServer()
	h := &handler.UserServer{Store: st}

	pb.RegisterUserServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
