package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("AUTH_DB_HOST"),
		os.Getenv("AUTH_DB_USER"),
		os.Getenv("AUTH_DB_PASSWORD"),
		os.Getenv("AUTH_DB_NAME"),
		os.Getenv("AUTH_DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Database connection error: ", err)
	} else {
		fmt.Println("Connected to the database")
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("failed to migrate db: ", err)
	}
}
