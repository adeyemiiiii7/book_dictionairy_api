package interfaces

import "example/go_api_tutorial/internal/models"

// UserRepository defines the contract for user data operations
type UserRepository interface {
	// Create operations
	Create(user *models.User) error
	
	// Read operations
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	
	// Update operations
	Update(user *models.User) error
	UpdatePassword(id uint, hashedPassword string) error
	
	// Delete operations
	Delete(id uint) error
	
	// Authentication helpers
	ExistsByUsername(username string) (bool, error)
	ExistsByEmail(email string) (bool, error)
}
