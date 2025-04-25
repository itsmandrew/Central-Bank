package main

import (
	"fmt"
	"log"

	"github.com/itsmandrew/Central-Bank/internal/config"
	"github.com/itsmandrew/Central-Bank/internal/db"
)

func main() {

	// Loading .env into os.Env
	cfg := config.Load()

	// Grab DB url
	dbURL := cfg.DBUrl
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set in the environment")
	}

	// Connect to Postgres database
	database := db.Connect(dbURL)
	defer database.Close()

	fmt.Println("✅ Successfully connected to Postgres!")
}
