package database

import (
	"log"
	"skrik/internal/config"
	entities "skrik/internal/entities/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDb() (*gorm.DB, error) {
	dsn :=
		"host=" + config.AppConfig.Postgres_host +
			"user=" + config.AppConfig.Postgres_user +
			"password=" + config.AppConfig.Postgres_password +
			"dbname=" + config.AppConfig.Postgres_db +
			"port=" + config.AppConfig.Postgres_port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalln("couldn't connect to the database. err: ", err)
	}
	err = db.AutoMigrate(entities.User{})
	if err != nil {
		log.Fatalln("couldn't migrate model(s). err: ", err)
	}
	return db, nil
}
