package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env.development")
	}
}
