package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomType int

const (
	SingleRoom RoomType = iota + 1
	DoubleRoom
	FamilyRoom
	DeluxeRoom
)

type Room struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomNO     int                `bson:"roomNo" json:"roomNo"`
	Type       RoomType           `bson:"type" json:"type"`
	Price      float64            `bson:"price" json:"price"`
	IsOccupied bool               `bson:"isOccupied" json:"isOccupied"`
	HotelID    primitive.ObjectID `bson:"hotelId,omitempty" json:"hotelId,omitempty"`
}
