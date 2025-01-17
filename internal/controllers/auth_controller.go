package controllers

import (
	"skrik/internal/entities"
	usecases "skrik/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	usecase *usecases.AuthUsecase
}

func NewAuthController(authUsecase *usecases.AuthUsecase, app *fiber.App) {
	controller := &AuthController{usecase: authUsecase}
	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return entities.NewBadRequestError("failed to read request body")
	}
	token, err := ac.usecase.Authorize(user.Username, user.Password)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString(token)
}
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	token, err := ac.usecase.Register(&user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString(token)
}

/* here's the refresh token generation functionality, but i can't properly test it so...
func (ac *AuthController) Refresh(c *fiber.Ctx) error{
	var request struct {
		Refresh_token string `json:"refresh_token"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	accessToken, err := ac.usecase.CompareRefreshTokens(request.Refresh_token)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "refresh tokens doesn't match"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": accessToken})
}*/
