package main

import (
	"log"
	"skrik/config"
	"skrik/database"
	_ "skrik/docs"
	"skrik/messaging"
	"skrik/routes"
	"skrik/signaling"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	EnvConfig := config.GetConfig()
	database.SetupDb()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081, http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))
	messaging.Test(app)
	signaling.VoiceHandler(app)
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))

}
