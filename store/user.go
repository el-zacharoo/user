package store

import (
	"context"
	"log"

	pb "github.com/el-zacharoo/user/user.v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"
)

type Storer interface {
	AddUser(u *pb.User, md metadata.MD) error
	QueryUsers(qr *pb.QueryRequest, md metadata.MD) ([]*pb.User, int64, error)
}

func (s Store) AddUser(u *pb.User, md metadata.MD) error {
	_, err := s.locaColl.InsertOne(context.Background(), u)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (s Store) QueryUsers(qr *pb.QueryRequest, md metadata.MD) ([]*pb.User, int64, error) {

	filter := bson.M{}

	if qr.Name != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"name": qr.Name},
		}}
	}

	opt := options.FindOptions{
		Skip:  &qr.Offset,
		Limit: &qr.Limit,
		Sort:  bson.M{"name": 1},
	}

	pg := pb.QueryResponse{User: []*pb.User{}}

	ctx := context.Background()
	cursor, err := s.locaColl.Find(ctx, filter, &opt)
	if err != nil {
		return nil, 0, err
	}

	docs := []pb.User{}
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, 0, err
	}

	pg.Matches, err = s.locaColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return nil, 0, err
}
