package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dapr/go-sdk/service/common"
	pbauth "github.com/el-zacharoo/user/internal/gen/auth/v1"
)

var sub = &common.Subscription{
	PubsubName: "user",
	Topic:      "adduser",
	Route:      "/Create",
}

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
