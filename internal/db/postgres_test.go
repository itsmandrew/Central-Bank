package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect_PingError(t *testing.T) {
	// Create the mock with ping monitoring enabled
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer mockDB.Close()

	// Store the original function and replace it safely
	origOpen := openDB
	openDB = func(driver, dsn string) (*sql.DB, error) {
		return mockDB, nil
	}
	// Make sure we restore the original function
	defer func() {
		openDB = origOpen
	}()

	// Set up ping expectation
	mock.ExpectPing().WillReturnError(fmt.Errorf("ping failed"))

	// Call the function under test
	conn, err := Connect("any-dsn")

	// Check results
	if err == nil {
		t.Error("Connect() should have returned an error")
	} else if err.Error() != "ping failed" {
		t.Errorf("Connect() error = %v, want %q", err, "ping failed")
	}

	// We expect conn to be nil when there's an error
	if conn != nil {
		conn.Close()
		t.Error("Connect() returned non-nil connection with error")
	}

	// Verify all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestConnect_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %s", err)
	}
	defer mockDB.Close()

	origOpen := openDB
	openDB = func(driver, dsn string) (*sql.DB, error) {
		return mockDB, nil
	}
	defer func() { openDB = origOpen }()

	mock.ExpectPing() // Expect ping to succeed

	conn, err := Connect("any-dsn")
	if err != nil {
		t.Fatalf("Connect() error = %v, want nil", err)
	}

	if conn == nil {
		t.Fatal("Connect() returned nil connection without error")
	}
	defer conn.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
