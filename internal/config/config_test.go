package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/itsmandrew/central-finance/internal/config"
)

func TestLoadFromEnv(t *testing.T) {

	// -- Arrange: set up known env variables --
	// Sets and automatically restores after the test
	t.Setenv("PORT", "1234")
	t.Setenv("DATABASE_URL", "postgres://foo:bar@localhost/baz")
	t.Setenv("PLAID_CLIENT_ID", "pid123")
	t.Setenv("PLAID_SECRET", "donut")
	t.Setenv("PLAID_ENV", "sandbox")

	// -- Act: call Load() --
	cfg := config.Load()

	// -- Assert: each field matches what we put in env --
	if got, want := cfg.Port, "1234"; got != want {
		t.Errorf("Port = %q; want %q", got, want)
	}

	if got, want := cfg.DBUrl, "postgres://foo:bar@localhost/baz"; got != want {
		t.Errorf("DBUrl = %q; want %q", got, want)
	}

	if got, want := cfg.PlaidClientID, "pid123"; got != want {
		t.Errorf("PlaidClientID = %q; want %q", got, want)
	}

	if got, want := cfg.PlaidSecret, "donut"; got != want {
		t.Errorf("PlaidSecret = %q; want %q", got, want)
	}

	if got, want := cfg.PlaidEnv, "sandbox"; got != want {
		t.Errorf("PlaidEnv = %q; want %q", got, want)
	}
}

func TestLoad_FromDotEnvFile(t *testing.T) {
	dir := t.TempDir()
	fmt.Println(dir)

	envContent := []byte(`
	PORT=9000
	DATABASE_URL=mysql://u:p@/dbname
	PLAID_CLIENT_ID=ci999
	PLAID_SECRET=sx888
	PLAID_ENV=development
	`)

	if err := os.WriteFile(filepath.Join(dir, ".env"), envContent, 0o600); err != nil {
		t.Fatalf("writing .env file: %v", err)
	}

	oldWd, err := os.Getwd()

	if err != nil {
		t.Fatalf("getting cwd: %v", err)
	}

	defer os.Chdir(oldWd)

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir to temp dir: %v", err)
	}

	cfg := config.Load()

	// -- Assert: values reflect what's in .env --
	if got, want := cfg.Port, "9000"; got != want {
		t.Errorf("Port = %q; want %q", got, want)
	}
	if got, want := cfg.DBUrl, "mysql://u:p@/dbname"; got != want {
		t.Errorf("DBUrl = %q; want %q", got, want)
	}
	if got, want := cfg.PlaidClientID, "ci999"; got != want {
		t.Errorf("PlaidClientID = %q; want %q", got, want)
	}
	if got, want := cfg.PlaidSecret, "sx888"; got != want {
		t.Errorf("PlaidSecret = %q; want %q", got, want)
	}
	if got, want := cfg.PlaidEnv, "development"; got != want {
		t.Errorf("PlaidEnv = %q; want %q", got, want)
	}
}
