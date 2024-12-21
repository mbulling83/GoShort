package db

import (
	"log"
)

// RunMigrations manually runs any pending database migrations
func RunMigrations() {
	if err := DB.Migrator().AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully.")
}
