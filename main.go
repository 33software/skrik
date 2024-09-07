package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Request struct {
	Username string `json:"name"`
	Mail     string `json:"email"`
}
type User struct {
	Userid int `json:"id"`
}

func main() {

	app := fiber.New()

	app.Get("/user", func(c fiber.Ctx) error {
		if userid := c.Query("userid"); userid != "" {
			return c.SendString("Hello " + userid)
		}
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/user", func(c fiber.Ctx) error {
		var request Request

		if err := c.Bind().Body(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.JSON(fiber.Map{
			"username": request.Username,
			"Email":    request.Mail,
		})
	})

	app.Put("/user", func(c fiber.Ctx) error {
		var request Request
		userid := c.Query("userid")
		if userid == "" {
			return c.Status(fiber.StatusBadRequest).SendString(fiber.ErrBadRequest.Error())
		}
		
		if err := c.Bind().Body(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		

		return c.JSON(
			fiber.Map{
				"userid":userid,
				"username":request.Username,
				"email":request.Mail,

			})
	})

	app.Delete("/user", func(c fiber.Ctx) error {
		userid := c.Query("userid")
		if userid == "" {
			return c.Status(fiber.StatusBadRequest).SendString(fiber.ErrBadRequest.Error())
		}
		return c.SendString("Delete " + userid)
	})

	log.Fatal(app.Listen(":3000"))

}
