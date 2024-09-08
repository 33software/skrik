package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type UserSchema struct {
	UserId   int    `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {

	app := fiber.New()

	app.Get("/user", func(c fiber.Ctx) error {
		if userid := c.Query("userid"); userid != "" {
			return c.Status(fiber.StatusOK).SendString("Hello " + userid)
		}
		return c.Status(fiber.StatusOK).SendString("Hello World")
	})

	app.Post("/user", func(c fiber.Ctx) error {
		var request UserSchema

		if err := c.Bind().Body(&request); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
		}
		response := fiber.Map{
			"username": request.Username,
			"email":    request.Email,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Put("/user", func(c fiber.Ctx) error {
		var request UserSchema
		userid := c.Query("userid")
		if userid == "" {
			return c.Status(fiber.StatusNotFound).SendString(fiber.ErrBadRequest.Error())
		}

		if err := c.Bind().Body(&request); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fiber.ErrInternalServerError.Error())
		}

		return c.JSON(
			fiber.Map{
				"userid":   request.UserId,
				"username": request.Username,
				"email":    request.Email,
			})
	})

	app.Delete("/user", func(c fiber.Ctx) error {
		userid := c.Query("userid")
		if userid == "" {
			return c.Status(fiber.StatusBadRequest).SendString(fiber.ErrBadRequest.Error())
		}
		return c.Status(fiber.StatusOK).SendString("Delete " + userid)
	})

	log.Fatal(app.Listen(":3000"))

}
