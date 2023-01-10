package config

import (
	"fmt"
	"os"
)

func RetrieveDbDns() string {
	dbHost := os.Getenv("PGHOST")
	dbName := os.Getenv("PGDATABASE")
	dbPass := os.Getenv("PGPASSWORD")
	dbPort := os.Getenv("PGPORT")
	dbUser := os.Getenv("PGUSER")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
