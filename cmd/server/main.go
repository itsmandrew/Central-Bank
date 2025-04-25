package main

import (
	"fmt"
	"log"

	"github.com/itsmandrew/Central-Bank/internal/config"
	"github.com/itsmandrew/Central-Bank/internal/db"
	"github.com/itsmandrew/Central-Bank/internal/transport"
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
	database, err := db.Connect(dbURL)

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		log.Fatalf("Database connection is required, exiting: %v", err)
	}

	defer database.Close()

	fmt.Println("✅ Successfully connected to Postgres!")

	router := transport.NewRouter(
		cfg, database,
	)

	log.Printf("Starting server on %s...", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}
