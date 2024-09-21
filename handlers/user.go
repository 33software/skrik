package handlers

import (
	jwtGen "audio-stream-golang/JWT"
	"audio-stream-golang/database"
	"audio-stream-golang/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetUser gets a user by ID
// @Summary Get user
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param userid query string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/users [get]
func GetUser(c *fiber.Ctx) error {
	var request models.User
	user := c.Query("userid")
	if user == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"})
	}

	if err := database.DataBase.First(&request, "ID= ?", user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

// @Summary Create a new user
// @Description Create a new user with the input data
// @Tags users
// @Accept  json
// @Produce  plain
// @Param user body models.User true "User data"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
	var request models.User
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"})
	}
	request.Password = string(hashedPassword)
	newUser := models.User{Email: request.Email, Username: request.Username, Password: request.Password}
	if err := database.DataBase.Create(&newUser).Error; err != nil {
		log.Println("couldn't create database record", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"})
	}

	token, err := jwtGen.GenerateJWT(newUser.ID)
	if err != nil {
		log.Println("couldn't create JWT token", err)
	}

	return c.Status(fiber.StatusOK).SendString(token)
}

// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param userid query string true "User ID"
// @Param user body models.User true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/users [put]
func UpdateUser(c *fiber.Ctx) error {
	var request map[string]interface{}
	var user models.User
	userid := c.Query("userid")
	if userid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"})
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"})
	}
	if err := database.DataBase.First(&user, "ID=?", userid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"})
	}

	for key, value := range request {
		if value == "" || value == nil {
			delete(request, key)
		}
	}

	if err := database.DataBase.Model(&user).Updates(request).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Error updating user"})
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param userid query string true "User ID"
// @Success 200 {string} string "User deleted"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/users [delete]
func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	userid := c.Query("userid")
	if userid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"})
	}

	if err := database.DataBase.Delete(&user, userid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).SendString("Delete " + userid)
}

