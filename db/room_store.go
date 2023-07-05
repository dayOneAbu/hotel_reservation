package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/dayoneabu/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	roomColl = "rooms"
)

type RoomStore interface {
	GetRoom(ctx context.Context, roomId string) (*types.Room, error)
	GetHotelRooms(ctx context.Context, hotelID string) ([]*types.Room, error)
	CreateRoom(ctx context.Context, hotelId string, room *types.Room) (*types.Room, error)
	UpdateRoom(ctx context.Context, roomId string, room *types.Room) error
	DeleteRoom(ctx context.Context, roomId string) error
}

type MongoRoomStore struct {
	Client     *mongo.Client
	Coll       *mongo.Collection
	HotelStore HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		Client:     client,
		Coll:       client.Database(dbName).Collection(roomColl),
		HotelStore: hotelStore,
	}
}
func (s *MongoRoomStore) GetRoom(ctx context.Context, roomId string) (*types.Room, error) {
	var room *types.Room
	oid, err := primitive.ObjectIDFromHex(roomId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	if err := s.Coll.FindOne(ctx, filter).Decode(&room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *MongoRoomStore) CreateRoom(ctx context.Context, hotelId string, room *types.Room) (*types.Room, error) {
	hotelOID, err := primitive.ObjectIDFromHex(hotelId)
	if err != nil {
		return nil, err
	}
	room.HotelID = hotelOID
	res, err := s.Coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)

	if err := s.HotelStore.UpdateHotelRoom(ctx, hotelId, room.ID); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *MongoRoomStore) GetHotelRooms(ctx context.Context, hotelID string) ([]*types.Room, error) {
	var rooms []*types.Room
	oid, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"hotelId": oid}
	cur, err := s.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}
func (s *MongoRoomStore) UpdateRoom(ctx context.Context, roomId string, room *types.Room) error {

	oid, err := primitive.ObjectIDFromHex(roomId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	// s.Coll.UpdateOne()
	if err := s.Coll.FindOneAndUpdate(ctx, filter, room).Decode(&room); err != nil {
		return err
	}
	fmt.Println(room)
	return nil

}
func (s *MongoRoomStore) DeleteRoom(ctx context.Context, roomId string) error {
	oid, err := primitive.ObjectIDFromHex(roomId)
	if err != nil {
		return err
	}
	deletedRoom, err := s.Coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil
	}
	if deletedRoom.DeletedCount != 0 {
		return nil
	}
	return errors.New("something went wrong please try again")

}

// TODO: populating data from object id
/*
// Sort by name in ascending order
	aggSort := bson.M{"$sort": bson.M{"roomNo": 1}}
	filter := bson.M{"$match": oid}
	// Populate Parent field
	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "Room",  // the collection name
		"localField":   "rooms", // the field on the child struct
		"foreignField": "_id",   // the field on the parent struct
		"as":           "rooms", // the field to populate into
	}}
	// qry := []bson.M{
	// 	{
	// 		"$match": filter,
	// 	},
	// 	{
	// 		"$lookup": bson.M{
	// 			// Define the tags collection for the join.
	// 			"from": "rooms",
	// 			// Specify the variable to use in the pipeline stage.
	// 			"let": bson.M{
	// 				"rooms": "$rooms",
	// 			},
	// 			"pipeline": []bson.M{
	// 				// Select only the relevant tags from the tags collection.
	// 				// Otherwise all the tags are selected.
	// 				// {
	// 				// 	"$match": bson.M{
	// 				// 		"$expr": bson.M{
	// 				// 			"$in": []interface{}{
	// 				// 				"$_id",
	// 				// 				"$$rooms",
	// 				// 			},
	// 				// 		},
	// 				// 	},
	// 				// },
	// 				// Sort tags by their roomNo field in asc. -1 = desc
	// 				{
	// 					"$sort": bson.M{
	// 						"roomNo": 1,
	// 					},
	// 				},
	// 			},
	// 			// Use tags as the field name to match struct field.
	// 			"as": "rooms",
	// 		},
	// 	},
	// }

	fmt.Println("hotel =", hotel)
	fmt.Println("hotel id =", hotelId)
	fmt.Println("hotel filter=", filter)

	// cur, err := s.Coll.Aggregate(ctx, []bson.M{
	// 	filter, aggSort, aggPopulate,
	// })
	// fmt.Println("hotel ==", cur)
	// if err != nil {
	// 	return nil, err
	// }
	// if err := cur.All(ctx, &hotel); err != nil {
	// 	return nil, err
	// }

*/
