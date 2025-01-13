package controllers

import (
	"skrik/internal/entities"
	usecases "skrik/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	usecase *usecases.AuthUsecase
}

func NewAuthController(authUsecase *usecases.AuthUsecase) *AuthController {
	return &AuthController{usecase: authUsecase}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	token, err := ac.usecase.Authorize(user.Username, user.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token:": token})
}
