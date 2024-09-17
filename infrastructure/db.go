package infrastructure

import (
	"fmt"
	"log"
	"weather-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "host=localhost dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.Weather{})
	if err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	fmt.Println("Database connection successful and migrations completed.")
	return db
}
