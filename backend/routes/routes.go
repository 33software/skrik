package routes

import (
	"skrik/handlers"

	jwtGen "skrik/JWT"
	_"skrik/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	account := api.Group(("/account"))
	account.Post("/register", handlers.CreateUser)
	account.Post("/login", handlers.Login)
	account.Get("/verify", handlers.VerifyEmail)
	account.Post("/reset", handlers.ResetPassword)
	account.Post("/resetendpoint", handlers.ResetEndpoint)

	account.Use(jwtGen.JwtProtected())
	account.Get("", handlers.GetUser)
	account.Put("", handlers.UpdateUser)
	account.Delete("/users", handlers.DeleteUser)
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

}
