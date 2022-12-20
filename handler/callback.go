package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	pbauth "github.com/el-zacharoo/user/internal/gen/auth/v1"
	pbsub "github.com/el-zacharoo/user/internal/gen/user/v1"
)

type CallbackServer struct {
	Server UserServer
	pb.UnimplementedAppCallbackServer
}

// Dapr will call this method to get the list of topics the app wants to subscribe to.
func (d CallbackServer) ListTopicSubscriptions(ctx context.Context, in *emptypb.Empty) (*pb.ListTopicSubscriptionsResponse, error) {

	return &pb.ListTopicSubscriptionsResponse{
		Subscriptions: []*pb.TopicSubscription{{
			PubsubName: "user",
			Topic:      "create",
			Routes:     &pb.TopicRoutes{Default: "/create"},
		}},
	}, nil
}

// OnTopicEvent is fired for events subscribed to.
// Dapr sends published messages in a CloudEvents 0.3 envelope.
func (d CallbackServer) OnTopicEvent(ctx context.Context, in *pb.TopicEventRequest) (*pb.TopicEventResponse, error) {
	var p pbsub.User
	if err := protojson.Unmarshal(in.Data, &p); err != nil {
		return &pb.TopicEventResponse{Status: pb.TopicEventResponse_DROP},
			status.Errorf(codes.Aborted, "issue unmarshalling data: %v", err)
	}

	switch in.Path {
	case "/create":
		createAuth(ctx, p.GivenName, p.FamilyName, p.Email, p.Id, p.Password)
	default:
		return &pb.TopicEventResponse{},
			status.Errorf(codes.Aborted, "unexpected path in OnTopicEvent: %s", in.Path)
	}

	return &pb.TopicEventResponse{}, nil
}

// Creates a new tenant on Auth0
func createAuth(ctx context.Context, givenName, familyName, email, password, id string) (*pbauth.CreateResponse, error) {

	url := os.Getenv("AUTH0_DOMAIN")
	key := os.Getenv("AUTH_CLIENT_ID")

	authUser := &pbauth.AuthUser{
		GivenName:  givenName,
		FamilyName: familyName,
		Email:      email,
		Password:   password,
		Connection: "Username-Password-Authentication",
		ClientId:   key,
		UserMetadata: &pbauth.UserMetadata{
			UserId: id,
		},
	}

	json_data, _ := json.Marshal(authUser)

	resp, _ := http.NewRequest("POST", "https://"+url+"/dbconnections/signup", bytes.NewBuffer(json_data))
	resp.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	_, err := client.Do(resp)
	if err != nil {
		return &pbauth.CreateResponse{}, err
	}
	return &pbauth.CreateResponse{AuthUser: authUser}, nil
}
