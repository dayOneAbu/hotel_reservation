package api

import (
	"github.com/dayoneabu/hotel_reservation/db"
	"github.com/dayoneabu/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandelGetUser(c *fiber.Ctx) error {
	user, err := h.userStore.GetUserByID(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(map[string]*types.User{"msg": user})
}

func (h *UserHandler) HandelGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetAllUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(map[string][]*types.User{"msg": users})
}
func (h *UserHandler) HandelPostUser(c *fiber.Ctx) error {
	// users, err := h.userStore.GetAllUsers(c.Context())
	// if err != nil {
	// 	return err
	// }
	firstName := c.Params("firstName")
	lastName := c.Params("lastName")
	var user = types.User{
		FirstName: firstName,
		LastName:  lastName,
	}
	h.userStore.CreateNewUser(c.Context(), &user)

	return nil
}
