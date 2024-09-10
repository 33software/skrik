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
	EnvConfig := config.GetConfig()
	dsn := "host=" + EnvConfig.Postgres_host +
		" user=" + EnvConfig.Postgres_user +
		" password=" + EnvConfig.Postgres_password +
		" dbname=" + EnvConfig.Postgres_db +
		" port=" + EnvConfig.Postgres_port
	Database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("couldn't connect to database", err)
		return
	}
	app := fiber.New()
	routes.SetupUserRoutes(app)
	routes.SetupSwagger(app)
	log.Fatal(app.Listen(EnvConfig.App_ip + ":" + EnvConfig.App_port))

	var result int64
	Database.Raw("SELECT 1123123123123123").Scan(&result)
	log.Println(result)
}
