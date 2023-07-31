package initializers

import (
	"codeles/library/data"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbError error

func Connect(connectionString string) {
	DB, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to DB")
}

func Migrate() {
	DB.AutoMigrate(&data.User{})
	log.Println("Database migrated")
}
