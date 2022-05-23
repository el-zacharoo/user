package handler

import (
	"context"
	"fmt"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	// event "go.buf.build/grpc/go/ticctech/common/event/v1"
	pbuser "github.com/el-zacharoo/user/gen/proto/go/user/v1"
)

type CallbackServer struct {
	UserServer UserServer
	pb.UnimplementedAppCallbackServer
}

// Dapr will call this method to get the list of topics the app wants to subscribe to.
func (d CallbackServer) ListTopicSubscriptions(ctx context.Context, in *emptypb.Empty) (*pb.ListTopicSubscriptionsResponse, error) {

	fmt.Println("ListTopicSubscriptions")
	return &pb.ListTopicSubscriptionsResponse{
		Subscriptions: []*pb.TopicSubscription{{
			PubsubName: "pubsub-test",
			Topic:      "mytopic",
			Routes: &pb.TopicRoutes{
				Rules: []*pb.TopicRule{
					{
						Match: `event.data.name == "Zach"`,
						Path:  "/update",
					},
				},
				Default: "/create",
			},
		}},
	}, nil
}

// OnTopicEvent is fired for events subscribed to.
// Dapr sends published messages in a CloudEvents 0.3 envelope.
func (d CallbackServer) OnTopicEvent(ctx context.Context, in *pb.TopicEventRequest) (*pb.TopicEventResponse, error) {

	fmt.Println("OnTopicEvent", in.Path, string(in.Data))
	// json event data -> event.EventData
	var user pbuser.User
	if err := protojson.Unmarshal(in.Data, &user); err != nil {
		return &pb.TopicEventResponse{Status: pb.TopicEventResponse_DROP},
			status.Errorf(codes.Aborted, "issue unmarshalling data: %v", err)
	}

	fmt.Println(&user)

	// // extract payload (google.protobuf.Any) from data
	// var pl event.IdentityPayload
	// if err := data.Payload.UnmarshalTo(&pl); err != nil {
	// 	return &pb.TopicEventResponse{Status: pb.TopicEventResponse_DROP},
	// 		status.Errorf(codes.Aborted, "issue unmarshalling payload: %v", err)
	// }

	switch in.Path {
	case "/create":
		// create checks for the identity
		// pl.Status == event.IdentityPayload_STATUS_CREATED
		// if err := d.CheckServer.addChecks(pl.IdentityId, data.TenantId, data.UserId); err != nil {
		// 	return &pb.TopicEventResponse{Status: pb.TopicEventResponse_DROP},
		// 		status.Errorf(codes.Aborted, "issue creating checks: %v", err)
		// }
	case "/update":
		// update the specified check
	default:
		return &pb.TopicEventResponse{},
			status.Errorf(codes.Aborted, "unexpected path in OnTopicEvent: %s", in.Path)
	}

	return &pb.TopicEventResponse{}, nil
}
