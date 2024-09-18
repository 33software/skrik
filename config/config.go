package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type CfgData struct {
	Postgres_user     string
	Postgres_db       string
	Postgres_password string
	Postgres_host     string
	Postgres_port     string
	App_ip            string
	App_port          string
	Jwt_keyword		  string

}

func GetConfig() CfgData {
	if err := godotenv.Load("dev.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	envConfig := CfgData{
		Postgres_user:     os.Getenv("POSTGRES_USER"),
		Postgres_db:       os.Getenv("POSTGRES_DB"),
		Postgres_password: os.Getenv("POSTGRES_PASSWORD"),
		Postgres_host:     os.Getenv("POSTGRES_HOST"),
		Postgres_port:     os.Getenv("POSTGRES_PORT"),
		App_ip:            os.Getenv("APP_IP"),
		App_port:          os.Getenv("APP_PORT"),
		Jwt_keyword:       os.Getenv("JWT_KEYWORD"),
	}
	return envConfig
}
