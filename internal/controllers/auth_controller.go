package controllers

import (
	"skrik/internal/entities"
	usecases "skrik/internal/usecases"
	"log"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	usecase *usecases.AuthUsecase
}

func NewAuthController(authUsecase *usecases.AuthUsecase, app *fiber.App){
	controller := &AuthController{usecase: authUsecase}
	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)
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

	return c.Status(fiber.StatusOK).SendString(token)
}
func (ac *AuthController) Register (c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		log.Println("failed to parse request. err: ", err)
		return err
	} 
	token, err := ac.usecase.Register(&user)
	if err != nil {
		log.Println("error! err: ", err)
	}

	return c.Status(fiber.StatusOK).SendString(token)
}
