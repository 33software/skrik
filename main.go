package main

import (
	"audio-stream-golang/config"
	"audio-stream-golang/db_connect"
	_ "audio-stream-golang/docs"
	"audio-stream-golang/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	EnvConfig := config.GetConfig()
	db_connect.SetupDb()
	app := fiber.New()
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))
}
