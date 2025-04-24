package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBUrl         string
	PlaidClientID string
	PlaidSecret   string
	PlaidEnv      string
}

func Load() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:          os.Getenv("PORT"),
		DBUrl:         os.Getenv("DATABASE_URL"),
		PlaidClientID: os.Getenv("PLAID_CLIENT_ID"),
		PlaidSecret:   os.Getenv("PLAID_SECRET"),
		PlaidEnv:      os.Getenv("PLAID_ENV"),
	}
}
