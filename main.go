package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Request struct {
	Username string `json:"name"`
	Mail     string `json:"email"`
}

/*type Params struct {
	Userid int
}*/

func main() {

	app := fiber.New()

	app.Get("/user", func(c fiber.Ctx) error {
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
		response := c.Queries()
		return c.JSON(response)
	})

	app.Delete("/user", func(c fiber.Ctx) error {

		return c.SendString("mb bby👋!")
	})

	log.Fatal(app.Listen(":3000"))

}
