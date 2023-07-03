package api

import (
	"errors"

	"github.com/dayoneabu/hotel_reservation/db"
	"github.com/dayoneabu/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

var params types.CreateUserParams

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandelGetUser(c *fiber.Ctx) error {
	user, err := h.userStore.GetUserByID(c.Context(), c.Params("id"))
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(map[string]string{"error": "user not Found"})
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
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	user, e := h.userStore.CreateNewUser(c.Context(), user)
	if e != nil {
		return e
	}
	return c.JSON(user)
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var update bson.M
	if err := c.BodyParser(&update); err != nil {
		return err
	}
	oid, oidErr := primitive.ObjectIDFromHex(c.Params("id"))
	if oidErr != nil {
		return oidErr
	}
	filter := bson.M{"_id": oid}
	update = bson.M{"$set": update}

	isUserUpdated, err := h.userStore.UpdateUser(c.Context(), filter, update)
	if err != nil {
		return err
	}
	if isUserUpdated {
		return c.JSON(map[string]string{"msg": "user updated successfully"})
	} else {
		return c.JSON(map[string]string{"msg": "please provide the right user"})
	}
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	if err := h.userStore.DeleteUser(c.Context(), c.Params("id")); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "user deleted successfully"})
}
