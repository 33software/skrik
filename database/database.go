package database

import (
	"audio-stream-golang/config"
	"audio-stream-golang/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() {
	EnvConfig := config.GetConfig()
	dsn := "host=" + EnvConfig.Postgres_host +
		" user=" + EnvConfig.Postgres_user +
		" password=" + EnvConfig.Postgres_password +
		" dbname=" + EnvConfig.Postgres_db +
		" port=" + EnvConfig.Postgres_port
	Database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("couldn't connect to database", err)
		return
	}
	err = Database.AutoMigrate(models.User{})
	if err != nil {
		log.Println("couldn't migrate the structure!", err)
		return
	}
	var result int64
	Database.Raw("SELECT 1").Scan(&result)
	log.Println(result)
}
