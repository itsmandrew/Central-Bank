package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvProvider interface {
	Get(key string) string
}

type osProvider struct{}

func (osProvider) Get(key string) string {
	return os.Getenv(key)
}

type Config struct {
	Port          string
	DBUrl         string
	PlaidClientID string
	PlaidSecret   string
	PlaidEnv      string
}

func LoadWith(p EnvProvider) *Config {
	cfg := &Config{
		Port:          p.Get("PORT"),
		DBUrl:         p.Get("DATABASE_URL"),
		PlaidClientID: p.Get("PLAID_CLIENT_ID"),
		PlaidSecret:   p.Get("PLAID_SECRET"),
		PlaidEnv:      p.Get("PLAID_ENV"),
	}

	// TODO: validate required fields here, log.Fatal if missing
	return cfg
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, relying on real env")
	}

	return LoadWith(osProvider{})
}
