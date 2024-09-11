package handlers

import (
	"github.com/gofiber/fiber/v2"
	"audio-stream-golang/models"
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
	if userid := c.Query("userid"); userid != "" {
		return c.Status(fiber.StatusOK).SendString("Hello " + userid)
	}
	return c.Status(fiber.StatusOK).SendString("Hello World")
}

// @Summary Create a new user
// @Description Create a new user with the input data
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body UserSchema true "User data"
// @Success 200 {object} UserSchema
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
	var request models.UserSchema

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
	}
	response := fiber.Map{
		"username": request.Username,
		"email":    request.Email,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body UserSchema true "Updated user data"
// @Success 200 {object} UserSchema
// @Router /api/users [put]
func UpdateUser(c *fiber.Ctx) error {
	var request models.UserSchema
	userid := c.Query("userid")
	if userid == "" {
		return c.Status(fiber.StatusNotFound).SendString(fiber.ErrBadRequest.Error())
	}

	// if err := c.Bind().Body(&request); err != nil {
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
	}

	return c.JSON(
		fiber.Map{
			"userid":   request.UserId,
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
