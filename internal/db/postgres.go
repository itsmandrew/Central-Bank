package db

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
)

// Make openDB a variable so it can be replaced in tests
var openDB = func(driver, dsn string) (*sql.DB, error) {
	return sql.Open(driver, dsn)
}

// Connect establishes a connection to the database using the provided URL
func Connect(dbURL string) (*sql.DB, error) {
	db, err := openDB("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close() // Close the connection before returning
		return nil, err
	}

	return db, nil
}
