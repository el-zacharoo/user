package store

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/el-zacharoo/user/internal/gen/user/v1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNotFound = errors.New("document not found")
)

type Storer interface {
	CreateUser(ctx context.Context, user *pb.User) error
	QueryUser(ctx context.Context, qr *pb.QueryRequest) ([]*pb.User, int64, error)
	GetUser(ctx context.Context, id string) (*pb.User, error)
	UpdateUser(id string, ctx context.Context, user *pb.User) error
	DeleteUser(id string) error
}

func (s Store) CreateUser(ctx context.Context, user *pb.User) error {
	_, err := s.locaColl.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (s Store) QueryUser(ctx context.Context, qr *pb.QueryRequest) ([]*pb.User, int64, error) {

	filter := bson.M{}
	if qr.SearchText != "" {
		filter = bson.M{"$text": bson.M{"$search": `"` + qr.SearchText + `"`}}
	}

	opt := options.FindOptions{
		Skip:  &qr.Offset,
		Limit: &qr.Limit,
		Sort:  bson.M{"date": -1},
	}

	cursor, err := s.locaColl.Find(ctx, filter, &opt)
	if err != nil {
		return nil, 0, err
	}

	var msgs []*pb.User
	if err := cursor.All(ctx, &msgs); err != nil {
		return nil, 0, err
	}

	matches, err := s.locaColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return msgs, matches, err
}

func (s Store) GetUser(ctx context.Context, id string) (*pb.User, error) {
	var user pb.User

	if err := s.locaColl.FindOne(context.Background(), bson.M{"id": id}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return &user, err
		}
		return &user, err
	}

	return &user, nil
}

func (s Store) UpdateUser(id string, ctx context.Context, user *pb.User) error {
	insertResult, err := s.locaColl.ReplaceOne(ctx, bson.M{"id": id}, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)

	return err
}

func (s Store) DeleteUser(id string) error {
	if _, err := s.locaColl.DeleteOne(context.Background(), bson.M{"id": id}); err != nil {
		return err
	}
	return nil
}
