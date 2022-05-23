package handler

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	// "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/el-zacharoo/user/gen/proto/go/user/v1"
	"github.com/el-zacharoo/user/store"
)

type UserServer struct {
	Store store.Storer
	Dapr  dapr.Client
	pb.UnimplementedUserServiceServer
}

func (u UserServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	user := req.User
	user.Id = uuid.NewString()
	// user.Created = timestamppb.Now()
	// user.Updated = timestamppb.Now()

	// publish event
	if err := u.Dapr.PublishEvent(
		context.Background(),
		"pubsub-test", "mytopic", user,
		dapr.PublishEventWithContentType("application/json"),
	); err != nil {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "error publishing event")
	}

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

	cur, mat, err := u.Store.QueryUsers(req, md)
	if err != nil {
		return &pb.QueryResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.QueryResponse{
		Cursor:  cur,
		Matches: mat,
	}, nil
}

func (u UserServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.GetResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	user, err := u.Store.GetUser(req.Id, md)
	if err != nil {
		return &pb.GetResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.GetResponse{User: user}, nil
}

func (u UserServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.UpdateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	user := req.User
	id := user.Id
	// user.Updated = timestamppb.Now()

	if err := u.Store.UpdateUser(id, md, user); err != nil {
		return &pb.UpdateResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.UpdateResponse{User: user}, nil
}

func (u UserServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.DeleteResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	if err := u.Store.DeleteUser(req.Id, md); err != nil {
		return &pb.DeleteResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.DeleteResponse{}, nil
}
