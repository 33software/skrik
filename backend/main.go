package main

import (
	"log"
	"skrik/config"
	"skrik/database"
	"skrik/handlers"
	"skrik/routes"
	_"skrik/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	EnvConfig := config.GetConfig()
	database.SetupDb()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))
	handlers.Setup(app)
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))

}
