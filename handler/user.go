package handler

import (
	"context"
	"fmt"

	"github.com/el-zacharoo/user/store"
	pb "github.com/el-zacharoo/user/user.v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Store interface {
	Create(qr *pb.CreateRequest, md metadata.MD)
}

type UserServer struct {
	Store store.Storer
	pb.UnimplementedUserServiceServer
}

func (s UserServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	fmt.Println("Write")

	user := req.User
	user.Id = uuid.NewString()

	if err := s.Store.AddUser(user, md); err != nil {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}
	return &pb.CreateResponse{User: user}, nil
}
