package db

import (
	"database/sql"
	"log"
)

func Connect(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Db open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Db ping error: %v", err)
	}

	return db
}
