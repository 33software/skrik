package routes

import (
	_ "audio-stream-golang/docs"
	"audio-stream-golang/handlers"

	jwtGen "audio-stream-golang/JWT"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/users", handlers.CreateUser)
	api.Post("/users/login", handlers.Login)

	api.Use(jwtGen.JwtProtected())
	api.Get("/users", handlers.GetUser)
	api.Put("/users", handlers.UpdateUser)
	api.Delete("/users", handlers.DeleteUser)
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

}
