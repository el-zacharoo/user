package handler

import (
	"context"
	"errors"
	"regexp"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"

	pb "github.com/el-zacharoo/user/internal/gen/user/v1"
	pbcnn "github.com/el-zacharoo/user/internal/gen/user/v1/userv1connect"
	"github.com/el-zacharoo/user/store"
)

type UserServer struct {
	// Dapr  dapr.Client
	Store store.Storer
	pbcnn.UnimplementedUserServiceHandler
}

func (s UserServer) Create(ctx context.Context, req *connect.Request[pb.CreateRequest]) (*connect.Response[pb.CreateResponse], error) {
	reqMsg := req.Msg
	user := reqMsg.User

	user.Id = uuid.NewString()

	// store functions
	if err := s.Store.CreateUser(ctx, user); err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	// response
	rsp := &pb.CreateResponse{User: user}
	return connect.NewResponse(rsp), nil
}

func (s UserServer) Query(ctx context.Context, req *connect.Request[pb.QueryRequest]) (*connect.Response[pb.QueryResponse], error) {
	reqUser := req.Msg

	if reqUser.SearchText != "" {
		pattern, err := regexp.Compile(`^[a-zA-Z@. ]+$`)
		if err != nil {
			return nil, connect.NewError(connect.CodeAborted, err)
		}
		if !pattern.MatchString(reqUser.SearchText) {
			return nil, connect.NewError(connect.CodeInvalidArgument,
				errors.New("invalid search text format"))
		}
	}

	cur, mat, err := s.Store.QueryUser(ctx, reqUser)
	if err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	rsp := &pb.QueryResponse{
		Cursor:  cur,
		Matches: mat,
	}
	return connect.NewResponse(rsp), nil
}

func (s UserServer) Get(ctx context.Context, req *connect.Request[pb.GetRequest]) (*connect.Response[pb.GetResponse], error) {
	reqUser := req.Msg.UserId

	// store functions
	user, err := s.Store.GetUser(ctx, reqUser)
	if err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	// response
	rsp := &pb.GetResponse{User: user}
	return connect.NewResponse(rsp), nil
}

func (s UserServer) Update(ctx context.Context, req *connect.Request[pb.UpdateRequest]) (*connect.Response[pb.UpdateResponse], error) {
	reqUser := req.Msg
	msg := reqUser.User

	// store functions
	if err := s.Store.UpdateUser(reqUser.UserId, ctx, msg); err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	// response
	rsp := &pb.UpdateResponse{User: msg}
	return connect.NewResponse(rsp), nil
}
