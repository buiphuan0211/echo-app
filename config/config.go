package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env ENV

func InitDotEnv() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	database := Database{URI: os.Getenv("DB_CONNECT_URI"), Name: os.Getenv("DATABASE_NAME")}
	jwt := Jwt{SecretKey: os.Getenv("JWT_SECRET_KEY"), TokenLife: os.Getenv("JWT_TOKEN_LIFE")}

	env = ENV{
		Port:     port,
		Database: database,
		Jwt:      jwt,
	}
}

func GetEnv() ENV {
	return env
}
