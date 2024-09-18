package handlers

import (
	"audio-stream-golang/database"
	"audio-stream-golang/models"
	"errors"
	"time"
	"log"
	"audio-stream-golang/config"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT (userid uint) (string, error) {
	EnvConfig := config.GetConfig()
	claims := jwt.MapClaims {
		"userid": userid,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return newToken.SignedString(EnvConfig.Jwt_keyword)
}

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
	user := c.Query("ID")
	
	if err := database.DataBase.First(&request, "ID= ?", user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("no such user")
		}
		return c.Status(fiber.StatusBadRequest).SendString("panic")
	}

	return c.Status(fiber.StatusOK).JSON(request)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
    }
    request.Password = string(hashedPassword)

	newUser := models.User{Email: request.Email, Username: request.Username, Password: request.Password}
	if err := database.DataBase.Create(&newUser).Error; err != nil{
		log.Println("couldn't create database record", err)
		return (fiber.ErrBadRequest)
	}

	token, err := GenerateJWT(newUser.ID)
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
