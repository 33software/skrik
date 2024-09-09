package main

import (
	"audio-stream-golang/config"
	_ "audio-stream-golang/docs"
	"audio-stream-golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	EnvConfig := config.GetConfig()
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))

}
