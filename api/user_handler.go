package api

import (
	"github.com/dayoneabu/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandelGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "john doe"})
}
func HandelGetUser(c *fiber.Ctx) error {
	james := types.User{
		ID:        "id001",
		FirstName: "john",
		LastName:  "doe",
	}
	return c.JSON(map[string]types.User{"msg": james})
}
