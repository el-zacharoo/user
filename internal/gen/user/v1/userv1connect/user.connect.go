// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/el-zacharoo/user/internal/gen/user/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	// create a message
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	// list messages
	Query(context.Context, *connect_go.Request[v1.QueryRequest]) (*connect_go.Response[v1.QueryResponse], error)
	// get a message by id
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
	// add more messages
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	// delete a message by id
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		create: connect_go.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+"/user.v1.UserService/Create",
			opts...,
		),
		query: connect_go.NewClient[v1.QueryRequest, v1.QueryResponse](
			httpClient,
			baseURL+"/user.v1.UserService/Query",
			opts...,
		),
		get: connect_go.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+"/user.v1.UserService/Get",
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+"/user.v1.UserService/Update",
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+"/user.v1.UserService/Delete",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	create *connect_go.Client[v1.CreateRequest, v1.CreateResponse]
	query  *connect_go.Client[v1.QueryRequest, v1.QueryResponse]
	get    *connect_go.Client[v1.GetRequest, v1.GetResponse]
	update *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Create calls user.v1.UserService.Create.
func (c *userServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Query calls user.v1.UserService.Query.
func (c *userServiceClient) Query(ctx context.Context, req *connect_go.Request[v1.QueryRequest]) (*connect_go.Response[v1.QueryResponse], error) {
	return c.query.CallUnary(ctx, req)
}

// Get calls user.v1.UserService.Get.
func (c *userServiceClient) Get(ctx context.Context, req *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Update calls user.v1.UserService.Update.
func (c *userServiceClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls user.v1.UserService.Delete.
func (c *userServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	// create a message
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	// list messages
	Query(context.Context, *connect_go.Request[v1.QueryRequest]) (*connect_go.Response[v1.QueryResponse], error)
	// get a message by id
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
	// add more messages
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	// delete a message by id
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/user.v1.UserService/Create", connect_go.NewUnaryHandler(
		"/user.v1.UserService/Create",
		svc.Create,
		opts...,
	))
	mux.Handle("/user.v1.UserService/Query", connect_go.NewUnaryHandler(
		"/user.v1.UserService/Query",
		svc.Query,
		opts...,
	))
	mux.Handle("/user.v1.UserService/Get", connect_go.NewUnaryHandler(
		"/user.v1.UserService/Get",
		svc.Get,
		opts...,
	))
	mux.Handle("/user.v1.UserService/Update", connect_go.NewUnaryHandler(
		"/user.v1.UserService/Update",
		svc.Update,
		opts...,
	))
	mux.Handle("/user.v1.UserService/Delete", connect_go.NewUnaryHandler(
		"/user.v1.UserService/Delete",
		svc.Delete,
		opts...,
	))
	return "/user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.Create is not implemented"))
}

func (UnimplementedUserServiceHandler) Query(context.Context, *connect_go.Request[v1.QueryRequest]) (*connect_go.Response[v1.QueryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.Query is not implemented"))
}

func (UnimplementedUserServiceHandler) Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.Get is not implemented"))
}

func (UnimplementedUserServiceHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.Update is not implemented"))
}

func (UnimplementedUserServiceHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.Delete is not implemented"))
}
