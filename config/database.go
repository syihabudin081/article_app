package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func DatabaseInit() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	var retries = 5
	for retries > 0 {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Database connected")
			return
		}
		retries--
		fmt.Println("Database connection failed, retrying...")
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	log.Fatalf("Failed to connect to the database after retries: %v", err)
}
