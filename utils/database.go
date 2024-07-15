package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=robin123 dbname=cryptodb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting in database", err)
		return nil
	}
	DB = db
}
