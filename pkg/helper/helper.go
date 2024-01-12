package helper

import (
	"github.com/gofiber/fiber/v2"
)

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJson(c *fiber.Ctx, code int, message string, data interface{}) error {
	json := MessageResponse{
		Message: message,
		Data:    data,
	}
	return c.Status(code).JSON(json)
}
