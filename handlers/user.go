package handlers

import (
	"audio-stream-golang/database"
	"audio-stream-golang/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUser gets a user by ID
// @Summary Get user
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param userid query string true "User ID"
// @Success 200 {string} string "Hello World"
// @Router /api/users [get]
func GetUser(c *fiber.Ctx) error {
	var request models.User
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
	}

	var user models.User
	if err := database.DataBase.First(&user, "email = ?", request.Email).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("no such user")
		}
	return c.Status(fiber.StatusBadRequest).SendString("panic")
	}

	return c.Status(fiber.StatusOK).JSON(user.Email)
}

// @Summary Create a new user
// @Description Create a new user with the input data
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserSchema true "User data"
// @Success 200 {object} models.UserSchema
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
	var request models.User

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
	}

	newUser := models.User{ID: request.ID, Email: request.Email, Username: request.Username, Password: request.Password}
	err := database.DataBase.Create(&newUser)
	if err != nil {
		log.Println("couldn't create database record", err)
		return (fiber.ErrBadRequest)
	}
	return c.Status(fiber.StatusOK).SendString("Success!")
}

// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.UserSchema true "Updated user data"
// @Success 200 {object} models.UserSchema
// @Router /api/users [put]
func UpdateUser(c *fiber.Ctx) error {
	var request models.User
	userid := c.Query("userid")
	if userid == "" {
		return c.Status(fiber.StatusNotFound).SendString(fiber.ErrBadRequest.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
	}

	return c.JSON(
		fiber.Map{
			"userid":   request.ID,
			"username": request.Username,
			"email":    request.Email,
		})
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted"
// @Router /api/users [delete]
func DeleteUser(c *fiber.Ctx) error {
	userid := c.Query("userid")
	if userid == "" {
		return c.Status(fiber.StatusBadRequest).SendString(fiber.ErrBadRequest.Error())
	}
	return c.Status(fiber.StatusOK).SendString("Delete " + userid)
}
