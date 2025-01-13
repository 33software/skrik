package controllers

import (
	"log"
	entities "skrik/internal/entities"
	usecases "skrik/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

// creating new user controller object
func NewUserController(usecase *usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (uc *UserController) RegisterUser(c *fiber.Ctx) {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		log.Println("failed to parse request. err: ", err)
		return
	}
	uc.usecase.RegisterUser(&user)

	log.Println("registered!", c.JSON(user))
}
func (uc *UserController) DeleteUser(id uint) {
	if id <= 0 {
		log.Println("incorrect id!")
		return
	}
	uc.usecase.DeleteUser(id)
}
