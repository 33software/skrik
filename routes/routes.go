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
	account := api.Group(("/account"))
	//account.Get("/verify", handlers.VerifyEmail)
	account.Post("/register", handlers.CreateUser)
	account.Post("/login", handlers.Login)

	account.Use(jwtGen.JwtProtected())
	account.Get("", handlers.GetUser)
	account.Put("", handlers.UpdateUser)
	account.Delete("/users", handlers.DeleteUser)
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

}
