package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"audio-stream-golang/routes"
)


func main() {

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
