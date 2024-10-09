package main

import (
	"skrik/config"
	"skrik/database"
	"skrik/routes"
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
