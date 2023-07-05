package api

import (
	"errors"

	"github.com/dayoneabu/hotel_reservation/db"
	"github.com/dayoneabu/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelHandler struct {
	hotelStore db.HotelStore
}

func NewHotelHandler(hotelStore db.HotelStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
	}
}

var hotel *types.Hotel

func (h *HotelHandler) HandelGetAllHotel(c *fiber.Ctx) error {
	users, err := h.hotelStore.GetAllHotel(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(map[string][]*types.Hotel{"msg": users})
}
func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	hotel, err := h.hotelStore.GetHotelByID(c.Context(), c.Params("id"))
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return c.JSON(map[string]string{"error": "hotel not Found"})
	}
	return c.JSON(map[string]*types.Hotel{"msg": hotel})
}
func (h *HotelHandler) HandelPostHotel(c *fiber.Ctx) error {
	if err := c.BodyParser(&hotel); err != nil {
		return err
	}
	hotel, err := h.hotelStore.CreateHotel(c.Context(), hotel)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}
func (h *HotelHandler) HandlePutHotel(c *fiber.Ctx) error {
	if err := c.BodyParser(&hotel); err != nil {
		return err
	}

	if err := h.hotelStore.UpdateHotel(c.Context(), c.Params("id"), hotel); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "hotel updated successfully"})
}

func (h *HotelHandler) HandleDeleteHotel(c *fiber.Ctx) error {
	if err := h.hotelStore.DeleteHotel(c.Context(), c.Params("id")); err != nil {
		return err
	}
	return c.JSON(map[string]string{"msg": "hotel removed successfully"})
}
