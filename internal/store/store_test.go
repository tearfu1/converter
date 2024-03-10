package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "example_user:example_user_password@tcp(localhost:3306)/example_db"
	}

	os.Exit(m.Run())
}
