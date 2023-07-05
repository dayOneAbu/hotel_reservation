package api

import (
	"github.com/dayoneabu/hotel_reservation/db"
	"github.com/dayoneabu/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	roomStore db.RoomStore
}

var room *types.Room

func NewRoomHandler(roomStore db.RoomStore) *RoomHandler {
	return &RoomHandler{
		roomStore: roomStore,
	}
}
func (h *RoomHandler) HandleGetRoom(c *fiber.Ctx) error {
	room, err := h.roomStore.GetRoom(c.Context(), c.Params("room_id"))
	if err != nil {
		return err
	}
	return c.JSON(room)
}

func (h *RoomHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	rooms, err := h.roomStore.GetHotelRooms(c.Context(), c.Params("hotel_id"))
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}
func (h *RoomHandler) HandlePostRoom(c *fiber.Ctx) error {
	if err := c.BodyParser(&room); err != nil {
		return err
	}
	room, err := h.roomStore.CreateRoom(c.Context(), c.Params("hotel_id"), room)
	if err != nil {
		return err
	}
	return c.JSON(room)
}

func (h *RoomHandler) HandlePutRoom(c *fiber.Ctx) error {
	if err := c.BodyParser(&room); err != nil {
		return err
	}
	if err := h.roomStore.UpdateRoom(c.Context(), c.Params("room_id"), room); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "room updated successfully"})
}
func (h *RoomHandler) HandleDeleteRoom(c *fiber.Ctx) error {
	if err := h.roomStore.DeleteRoom(c.Context(), c.Params("room_id")); err != nil {
		return err
	}

	return c.JSON(map[string]string{"msg": "room deleted successfully"})
}
