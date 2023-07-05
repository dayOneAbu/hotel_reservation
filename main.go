package main

import (
	"flag"

	"github.com/dayoneabu/hotel_reservation/api"
	"github.com/dayoneabu/hotel_reservation/db"

	"github.com/gofiber/fiber/v2"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAdd := flag.String("listenAdd", ":3000", "DB listening port")
	flag.Parse()
	client := db.ConnectTOMongo()
	app := fiber.New(config)
	apiV1 := app.Group("/api/v1")

	// handlers initialization
	var (
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)
		userStore  = db.NewMongoUserStore(client)

		hotelHandler = api.NewHotelHandler(hotelStore)
		roomHandler  = api.NewRoomHandler(roomStore)
		userHandler  = api.NewUserHandler(userStore)
	)

	app.Get("/", handleHome)
	// userHandlers
	apiV1.Get("/users", userHandler.HandelGetUsers)
	apiV1.Get("/users/:id", userHandler.HandelGetUser)
	apiV1.Post("/users/", userHandler.HandelPostUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	// hotelHandlers
	apiV1.Get("/hotels", hotelHandler.HandelGetAllHotel)
	apiV1.Get("/hotels/:id", hotelHandler.HandleGetHotelByID)
	apiV1.Post("/hotels", hotelHandler.HandelPostHotel)
	apiV1.Put("/hotels/:id", hotelHandler.HandlePutHotel)
	apiV1.Delete("/hotels/:hotelId/:roomId", hotelHandler.HandleDeleteHotel)
	// hotelHandlers
	apiV1.Get("/rooms", roomHandler.HandleGetRoom)
	apiV1.Get("/rooms/:hotel_id", roomHandler.HandleGetHotelRooms)
	apiV1.Post("/rooms", roomHandler.HandlePostRoom)
	apiV1.Put("/rooms/:room_id", roomHandler.HandlePutRoom)
	apiV1.Delete("/rooms/:room_id", roomHandler.HandleDeleteRoom)
	app.Listen(*listenAdd)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Hello, World!"})
}
