package main

import (
	"context"
	"fmt"

	"github.com/dayoneabu/hotel_reservation/db"
	"github.com/dayoneabu/hotel_reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	hotel := types.Hotel{
		Name:     "9to5 luxury hotel",
		Location: "22nd street, A.A",
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []*types.Room{
		{
			RoomNO:     10,
			Type:       types.SingleRoom,
			Price:      99.99,
			IsOccupied: false,
		},
		{
			RoomNO:     20,
			Type:       types.DoubleRoom,
			Price:      99.99,
			IsOccupied: false,
		},
		{
			RoomNO:     30,
			Type:       types.FamilyRoom,
			Price:      99.99,
			IsOccupied: false,
		},
		{
			RoomNO:     40,
			Type:       types.DeluxeRoom,
			Price:      99.99,
			IsOccupied: false,
		},
	}

	fmt.Println("DB Seeding Started...")
	fmt.Println("creating hotel collection")
	client := db.ConnectTOMongo()
	dbName := db.LoadEnv("DBNAME")
	if err := client.Database(dbName).Drop(context.TODO()); err != nil {
		fmt.Println(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	newHotel, err := hotelStore.CreateHotel(context.Background(), &hotel)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hotel collection created", newHotel.Name)

	for _, room := range rooms {
		newRoom, err := roomStore.CreateRoom(context.Background(), newHotel.ID.Hex(), room)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(newRoom)
	}
}
