package handler

import (
	"context"

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

func (u UserServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	user := req.User
	user.Id = uuid.NewString()

	if err := u.Store.AddUser(user, md); err != nil {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}
	return &pb.CreateResponse{User: user}, nil
}

func (u UserServer) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.QueryResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	usr, mat, err := u.Store.QueryUsers(req, md)
	if err != nil {
		return &pb.QueryResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.QueryResponse{
		User:    usr,
		Matches: mat,
	}, nil
}
