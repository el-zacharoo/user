package store

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type Storer[T proto.Message] interface {
	Create(msg, md metadata.MD) error
	// List(filter bson.M, md metadata.MD, opts ...ListOption) ([]T, int64, error)
	// Update(msg T, id string, md metadata.MD) error
	// Get(id string, md metadata.MD) (T, error)
	// Delete(id string, md metadata.MD, opts ...DeleteOption) error
}

func (s Store) Create(msg, md metadata.MD) error {
	_, err := s.locaColl.InsertOne(context.Background(), msg)
	return err
}
