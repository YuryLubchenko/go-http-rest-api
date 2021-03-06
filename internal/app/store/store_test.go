package store_test

import (
	"os"
	"testing")


var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable user=postgres password=123456"
	}

	os.Exit(m.Run())
}