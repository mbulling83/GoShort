package db

import (
	"GoShort/internal/models"
	"GoShort/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	dsn := config.Get("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate the database schema
	if err := DB.AutoMigrate(&models.URL{}, &models.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	log.Println("Database connection initialized successfully.")
}
