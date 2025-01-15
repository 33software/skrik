package controllers

import (
	"skrik/internal/entities"

	"github.com/gofiber/fiber/v2"
)

func ErrHandlerMiddleware(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {
		if appErr, ok := err.(*entities.AppErr); ok {
			return c.Status(appErr.Code).JSON(appErr.Message)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	return nil
}
