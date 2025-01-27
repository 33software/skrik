package app

import (
	"log"
	_ "skrik/docs"
	"skrik/internal/config"
	controllers "skrik/internal/controllers"
	"skrik/internal/database"
	repository "skrik/internal/repository"
	"skrik/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Run() {
	//loading .env configuration; connecting to the database
	config.LoadCfg()
	db, err := database.StartDb()
	if err != nil {
		log.Fatalln("couldn't start database. err: ", err)
	}
	app := fiber.New()
	app.Use("/", controllers.ErrHandlerMiddleware)
	app.Get("/swagger/*", swagger.HandlerDefault)

	chatRepo := repository.NewChatRepository(db)
	chatUsecase := usecases.NewChatUsecase(chatRepo)
	controllers.NewChatController(chatUsecase, app)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	controllers.NewUserController(userUsecase, app)
	authUsecase := usecases.NewAuthUsecase(userRepo)
	controllers.NewAuthController(authUsecase, app)

	app.Listen(config.AppConfig.App_ip + ":" + config.AppConfig.App_port)
}
