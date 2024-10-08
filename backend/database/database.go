package database

import (
	"skrik/config"
	"skrik/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func SetupDb() {
	EnvConfig := config.GetConfig()
	dsn := "host=" + EnvConfig.Postgres_host +
		" user=" + EnvConfig.Postgres_user +
		" password=" + EnvConfig.Postgres_password +
		" dbname=" + EnvConfig.Postgres_db +
		" port=" + EnvConfig.Postgres_port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Println("couldn't connect to database", err)
		return
	}
	DataBase = db
	err = DataBase.AutoMigrate(models.User{})
	if err != nil {
		log.Println("couldn't migrate database user model", err)
		return
	}

	var result int64
	DataBase.Raw("SELECT 1").Scan(&result)
	log.Println(result)
}
