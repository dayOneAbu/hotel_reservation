package db

import (
	"context"
	"errors"

	"github.com/dayoneabu/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	hotelColl = "hotels"
)

type HotelStore interface {
	GetAllHotel(ctx context.Context) ([]*types.Hotel, error)
	GetHotelByID(ctx context.Context, hotelId string) (*types.Hotel, error)
	CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, hotelId string, hotel *types.Hotel) error
	UpdateHotelRoom(ctx context.Context, hotelId string, roomId primitive.ObjectID) error
	DeleteHotel(ctx context.Context, hotelId string) error
}

type MongoHotelStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		Client: client,
		Coll:   client.Database(dbName).Collection(hotelColl),
	}
}

func (s *MongoHotelStore) GetAllHotel(ctx context.Context) ([]*types.Hotel, error) {
	var hotels []*types.Hotel
	res, err := s.Coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil

}
func (s *MongoHotelStore) GetHotelByID(ctx context.Context, hotelId string) (*types.Hotel, error) {
	var hotel *types.Hotel
	oid, err := primitive.ObjectIDFromHex(hotelId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	if err := s.Coll.FindOne(ctx, filter).Decode(&hotel); err != nil {
		return nil, err
	}
	return hotel, nil
}

func (s *MongoHotelStore) CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := s.Coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, hotelId string, hotel *types.Hotel) error {

	oid, oidErr := primitive.ObjectIDFromHex(hotelId)
	if oidErr != nil {
		return oidErr
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$set": hotel}
	updatedHotel, err := s.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updatedHotel.ModifiedCount != 0 {
		return nil
	}
	return errors.New("something went south please try again")
}
func (s *MongoHotelStore) UpdateHotelRoom(ctx context.Context, hotelId string, roomId primitive.ObjectID) error {

	hotel_id, oidErr := primitive.ObjectIDFromHex(hotelId)
	if oidErr != nil {
		return oidErr
	}
	filter := bson.M{"_id": hotel_id}
	update := bson.M{"$push": bson.M{"rooms": roomId}}
	updatedHotel, err := s.Coll.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return err
	}
	if updatedHotel.ModifiedCount != 0 {
		return nil
	}
	return errors.New("something went south please try again")
}
func (s *MongoHotelStore) DeleteHotel(ctx context.Context, hotelId string) error {
	oid, oidErr := primitive.ObjectIDFromHex(hotelId)
	if oidErr != nil {
		return oidErr
	}

	filter := bson.M{"_id": oid}
	deletedHotel, err := s.Coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if deletedHotel.DeletedCount != 0 {
		return nil
	}
	return errors.New("something went wrong please try again")
}
