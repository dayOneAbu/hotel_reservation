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

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	app.Get("/", handleHome)
	apiV1.Get("/users", userHandler.HandelGetUsers)
	apiV1.Get("/users/:id", userHandler.HandelGetUser)
	apiV1.Post("/users/", userHandler.HandelPostUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	app.Listen(*listenAdd)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Hello, World!"})
}
