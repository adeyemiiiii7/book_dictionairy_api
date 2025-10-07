package database

import (
	"example/go_api_tutorial/internal/models"
	"log"
)

// Migrate runs database migrations
func Migrate() error {
	log.Println("Running database migrations...")
	
	// Auto-migrate the schema
	err := DB.AutoMigrate(
		&models.User{},
		&models.Book{},
	)
	
	if err != nil {
		return err
	}
	
	log.Println("Database migrations completed successfully!")
	return nil
}
