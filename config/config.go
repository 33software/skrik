package config

import (
	"os"

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
}

func GetConfig() CfgData {
	godotenv.Load()
	envConfig := CfgData{
		Postgres_user:     os.Getenv("POSTGRES_USER"),
		Postgres_db:       os.Getenv("POSTGRES_DB"),
		Postgres_password: os.Getenv("POSTGRES_PASSWORD"),
		Postgres_host:     os.Getenv("POSTGRES_HOST"),
		Postgres_port:     os.Getenv("POSTGRES_PORT"),
		App_ip:            os.Getenv("APP_IP"),
		App_port:          os.Getenv("APP_PORT"),
	}
	return envConfig
}
