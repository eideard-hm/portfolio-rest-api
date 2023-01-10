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
	dbName := os.Getenv("db_name")
	host := os.Getenv("db_host")
	password := os.Getenv("db_pass")
	user := os.Getenv("db_user")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbName, password)

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
