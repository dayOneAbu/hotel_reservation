package main

import (
	"flag"

	"github.com/dayoneabu/hotel_reservation/api"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	name    string
	age     int
	address string
}

func main() {
	listenAdd := flag.String("listenAdd", ":3000", "DB listening port")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	app.Get("/", handleHome)
	apiV1.Get("/user", api.HandelGetUser)
	app.Listen(*listenAdd)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Hello, World!"})
}
