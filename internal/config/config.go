package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres_user     string
	Postgres_db       string
	Postgres_password string
	Postgres_host     string
	Postgres_port     string
	App_ip            string
	App_port          string
	Jwt_keyword       string
	Smtp_host         string
	Smtp_port         string
	Smtp_sender       string
}

var AppConfig Config

func LoadCfg() {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Fatalln("failed to read .env file. err: ", err)
	}

	AppConfig = Config{
		Postgres_user:     os.Getenv("POSTGRES_USER"),
		Postgres_db:       os.Getenv("POSTGRES_DB"),
		Postgres_password: os.Getenv("POSTGRES_PASSWORD"),
		Postgres_host:     os.Getenv("POSTGRES_HOST"),
		Postgres_port:     os.Getenv("POSTGRES_PORT"),
		App_ip:            os.Getenv("APP_IP"),
		App_port:          os.Getenv("APP_PORT"),
		Jwt_keyword:       os.Getenv("JWT_KEYWORD"),
		Smtp_sender:       os.Getenv("SMTP_SENDER"),
		Smtp_host:         os.Getenv("SMTP_HOST"),
		Smtp_port:         os.Getenv("SMTP_PORT"),
	}
}
