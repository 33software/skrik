package main

import (
	"log"
	"skrik/internal/config"
	controllers "skrik/internal/controllers/user"
	"skrik/internal/database"
	repository "skrik/internal/repository/user"
	usecases "skrik/internal/usecases/user"
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
}
