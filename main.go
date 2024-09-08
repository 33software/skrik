package main

import (
	"audio-stream-golang/config"
	"audio-stream-golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	EnvConfig := config.GetConfig()
	routes.SetupUserRoutes(app)
	log.Fatal(app.Listen(EnvConfig.App_port))

}
