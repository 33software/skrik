package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Request struct {
	username string `json:"name"`
	mail     string `json:"email"`
}

func main() {

	app := fiber.New()

	app.Get("/", func(d fiber.Ctx) error {
		return d.SendString("Hello, World 👋!")
	})

	app.Post("/register", func(d fiber.Ctx) error {
		var request Request

		if err := d.Bind().Body(&request); err != nil {
			return d.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return d.JSON(fiber.Map{
			"message": "урааА",
			"data":    request,
		})
	})

	app.Put("/", func(d fiber.Ctx) error {

		return d.SendString("ayy 33 got on the top 0_o👋!")
	})

	app.Delete("/", func(d fiber.Ctx) error {

		return d.SendString("mb bby👋!")
	})

	log.Fatal(app.Listen(":3000"))

}
