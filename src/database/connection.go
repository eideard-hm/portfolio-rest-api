package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConnection() {
	dbHost := os.Getenv("PGHOST")
	dbName := os.Getenv("PGDATABASE")
	dbPass := os.Getenv("PGPASSWORD")
	dbPort := os.Getenv("PGPORT")
	dbUser := os.Getenv("PGUSER")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	conn, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		db = conn
		log.Println("Connected successfull to database")
	}
}

// returns a handle to the DB object
func GetDb() *gorm.DB {
	return db
}
