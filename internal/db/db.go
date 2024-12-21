package db

import (
	"GoShort/internal/models"
	"GoShort/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and applies migrations
func InitDB() {
	// Load the database connection string from environment variables
	dsn := config.Get("DATABASE_URL")

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Apply migrations
	log.Println("Running migrations...")
	if err := DB.AutoMigrate(&models.URL{}); err != nil {
		log.Fatalf("Failed to migrate URL schema: %v", err)
	}
	log.Println("Migrations completed successfully.")

	log.Println("Database connection initialized successfully.")
}
