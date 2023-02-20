package database

import (
	"log"

	"github.com/eideard-hm/portfolio-rest-api/src/config"
	"github.com/eideard-hm/portfolio-rest-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConnection() {
	dsn := config.RetrieveDbDns()

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
func Migrate() {
	db.Debug().AutoMigrate(&models.User{})
}
