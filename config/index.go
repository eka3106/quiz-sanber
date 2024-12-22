package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	SecretJwt   string
}

var VarConfig Config

func init() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	VarConfig = Config{
		DB_USERNAME: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		SecretJwt:   os.Getenv("SECRET_JWT"),
	}
}
