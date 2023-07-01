package db

import (
	"context"

	"github.com/dayoneabu/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName   = "hotel_reservation"
	userColl = "users"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetAllUsers(context.Context) ([]*types.User, error)
	CreateNewUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbName).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetAllUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
func (s *MongoUserStore) CreateNewUser(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, nil
}
