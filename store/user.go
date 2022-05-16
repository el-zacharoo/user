package store

import (
	"context"
	"fmt"
	"log"

	pb "github.com/el-zacharoo/user/user.v1"
	"google.golang.org/grpc/metadata"
)

type Storer interface {
	AddUser(u *pb.User, md metadata.MD) error
	// List(filter bson.M, md metadata.MD, opts ...ListOption) ([]T, int64, error)
	// Update(msg T, id string, md metadata.MD) error
	// Get(id string, md metadata.MD) (T, error)
	// Delete(id string, md metadata.MD, opts ...DeleteOption) error
}

func (s Store) AddUser(u *pb.User, md metadata.MD) error {
	insertResult, err := s.locaColl.InsertOne(context.Background(), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
	return err
}
