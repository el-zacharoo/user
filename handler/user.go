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

type Server struct {
	pb.UnimplementedUserServiceServer
	Store store.Store
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	req.User.Id = uuid.NewString()

	if err := s.Store.Create(md, nil); err != nil {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}
	return &pb.CreateResponse{}, nil
}
