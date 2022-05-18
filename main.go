package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/el-zacharoo/user/handler"
	"github.com/el-zacharoo/user/store"

	pb "github.com/el-zacharoo/user/gen/proto/go/user/v1"
)

func main() {
	grpcSrv := grpc.NewServer()
	defer grpcSrv.Stop()         // stop server on exit
	reflection.Register(grpcSrv) // for postman

	h := &handler.UserServer{Store: store.Connect()}
	pb.RegisterUserServiceServer(grpcSrv, h)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
