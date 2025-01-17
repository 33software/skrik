package controllers

import (
	"skrik/internal/auth"
	entities "skrik/internal/entities"
	usecases "skrik/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

// creating new user controller object and registering api routes
func NewUserController(usecase *usecases.UserUsecase, app *fiber.App) {
	handler := &UserController{usecase: usecase}
	api := app.Group("/api")
	api.Use(auth.Middleware())
	api.Post("/deleteuser", handler.DeleteUser)
	api.Get("/me", handler.GetProfile)

}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return entities.NewBadRequestError("failed to read request body")
	}
	if user.ID <= 0 {
		return entities.NewBadRequestError("incorrect userid. debug: ")
	}
	uc.usecase.DeleteUser(user.ID)
	return c.SendStatus(fiber.StatusOK)
}
func (uc *UserController) GetProfile(c *fiber.Ctx) error {
	userid := c.Locals("userid").(int)
	user, err := uc.usecase.GetUserByID(uint(userid))
	if err != nil {
		return err
	}

	return c.JSON(user)
}

/* there might be a password change functionality idk
func (uc *UserController) ChangePassword (c *fiber.Ctx){

}*/
