package config

import "testing"

type MapProvider map[string]string

func (m MapProvider) Get(key string) string {
	return m[key]
}

// Easier to test, don't need to setup env variables, easier to scale.
// SSOT: production loader is the same and tests never touch the real .env files
func TestLoadWith(t *testing.T) {
	mp := MapProvider{
		"PORT":            "1234",
		"DATABASE_URL":    "postgres://foo:bar@localhost/baz",
		"PLAID_CLIENT_ID": "pid123",
		"PLAID_SECRET":    "sec456",
		"PLAID_ENV":       "sandbox",
	}

	cfg := LoadWith(mp)

	tests := []struct {
		name, got, want string
	}{
		{"Port", cfg.Port, "1234"},
		{"DBUrl", cfg.DBUrl, "postgres://foo:bar@localhost/baz"},
		{"PlaidClientID", cfg.PlaidClientID, "pid123"},
		{"PlaidSecret", cfg.PlaidSecret, "sec456"},
		{"PlaidEnv", cfg.PlaidEnv, "sandbox"},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("%s = %q; want %q", tt.name, tt.got, tt.want)
		}
	}
}
