package routes

import (
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"

	"audio-stream-golang/handlers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    // Пример маршрута для пользователей
    api.Get("/users", handlers.GetUser)
    api.Post("/users", handlers.CreateUser)
	api.Put("/users", handlers.UpdateUser)
	api.Delete("/users", handlers.DeleteUser)
}

func SetupSwagger(app *fiber.App){
	app.Get("/swagger/*", swagger.HandlerDefault)
}