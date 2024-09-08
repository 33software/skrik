package routes

import (
	_ "audio-stream-golang/docs"
	"audio-stream-golang/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", handlers.GetUser)
	api.Post("/users", handlers.CreateUser)
	api.Put("/users", handlers.UpdateUser)
	api.Delete("/users", handlers.DeleteUser)
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
