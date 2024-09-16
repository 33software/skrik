package main

import (
	"audio-stream-golang/config"
	"audio-stream-golang/database"
	_ "audio-stream-golang/docs"
	"audio-stream-golang/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	EnvConfig := config.GetConfig()
	database.SetupDb()
	app := fiber.New()
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))

}
