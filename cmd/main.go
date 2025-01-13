package main

import (
	"log"
	"skrik/internal/auth"
	"skrik/internal/config"
	controllers "skrik/internal/controllers"
	"skrik/internal/database"
	repository "skrik/internal/repository"
	"skrik/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

// this probably should be moved to app.go
func main() {
	//loading .env configuration; connecting to the database
	config.LoadCfg()
	db, err := database.StartDb()
	if err != nil {
		log.Fatalln("couldn't start database. err: ", err)
	}
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)
	authUsecase := usecases.NewAuthUsecase(userRepo)
	authController := controllers.NewAuthController(authUsecase)

	app := fiber.New()
	app.Post("/auth/login", authController.Login)

	app.Use("/api", auth.Middleware())
	//app.Post("api/delete", userController.DeleteUser())
}
