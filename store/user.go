package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"

	pb "github.com/el-zacharoo/user/user.v1"
)

type Storer interface {
	AddUser(u *pb.User, md metadata.MD) error
	QueryUsers(qr *pb.QueryRequest, md metadata.MD) ([]*pb.User, int64, error)
	GetUser(id string, md metadata.MD) (*pb.User, error)
	UpdateUser(id string, md metadata.MD, u *pb.User) error
}

func (s Store) AddUser(u *pb.User, md metadata.MD) error {
	_, err := s.locaColl.InsertOne(context.Background(), u)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (s Store) QueryUsers(qr *pb.QueryRequest, md metadata.MD) ([]*pb.User, int64, error) {
	var filter bson.M

	if qr.Name != "" {
		filter = bson.M{"$text": bson.M{"$search": `"` + qr.Name + `"`}}
	}

	opt := options.FindOptions{
		Skip:  &qr.Offset,
		Limit: &qr.Limit,
		Sort:  bson.M{"name": -1},
	}

	ctx := context.Background()
	cursor, err := s.locaColl.Find(ctx, filter, &opt)
	if err != nil {
		return nil, 0, err
	}

	var users []*pb.User
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, 0, err
	}

	matches, err := s.locaColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return users, matches, err
}

func (s Store) GetUser(id string, md metadata.MD) (*pb.User, error) {
	var u pb.User

	if err := s.locaColl.FindOne(context.Background(), bson.M{"id": id}).Decode(&u); err != nil {
		if err == mongo.ErrNoDocuments {
			return &u, err
		}
		return &u, err
	}

	return &u, nil
}

func (s Store) UpdateUser(id string, md metadata.MD, u *pb.User) error {
	insertResult, err := s.locaColl.ReplaceOne(context.Background(), bson.M{"id": id}, u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)

	return err
}
