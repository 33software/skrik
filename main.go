package main

import (
	"audio-stream-golang/config"
	_ "audio-stream-golang/docs"
	"audio-stream-golang/routes"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	EnvConfig := config.GetConfig()
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))

	dsn := "host=" + EnvConfig.Postgres_host +
		" user=" + EnvConfig.Postgres_user +
		" password=" + EnvConfig.Postgres_password +
		" dbname=" + EnvConfig.Postgres_db +
		" dbPort=" + EnvConfig.Postgres_port

	gorm.Open(postgres.Open(dsn))
}
