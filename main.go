package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"audio-stream-golang/routes"

	_ "audio-stream-golang/docs"
)

func main() {

	app := fiber.New()

	routes.SetupUserRoutes(app)

	routes.SetupSwagger(app)

	log.Fatal(app.Listen("0.0.0.0:8080"))

}
