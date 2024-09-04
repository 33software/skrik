package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(d fiber.Ctx) error {
		// Send a string response to the client
		return d.SendString("Hello, World 👋!")
	})

	app.Post("/", func(d fiber.Ctx) error {

		return d.SendString(" Hello, POST👋!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
