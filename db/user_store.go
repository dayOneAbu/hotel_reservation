package db

import (
	"context"
	"errors"

	"github.com/dayoneabu/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName   = "hotel_reservation"
	userColl = "users"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetAllUsers(context.Context) ([]*types.User, error)
	CreateNewUser(context.Context, *types.User) (*types.User, error)
	UpdateUser(ctx context.Context, id string, user *types.User) (bool, error)
	DeleteUser(context.Context, string) error
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
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, oidErr := primitive.ObjectIDFromHex(id)
	if oidErr != nil {
		return oidErr
	}
	_, err := s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	// TODO: implement some sort of verification
	// fmt.Println(res.DeletedCount != 0)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, id string, user *types.User) (bool, error) {
	opts := options.Update().SetUpsert(false)

	oid, oidErr := primitive.ObjectIDFromHex(id)
	if oidErr != nil {
		return false, oidErr
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": user}
	updatedUser, err := s.coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return false, err
	}
	if updatedUser.ModifiedCount != 0 {
		return true, nil
	}
	return false, errors.New("something went south please try again")
}
