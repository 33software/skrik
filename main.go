package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"audio-stream-golang/routes"
)

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen("0.0.0.0:8080"))

}
