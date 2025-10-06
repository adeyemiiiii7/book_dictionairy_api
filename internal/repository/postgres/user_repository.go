package postgres

import (
	"example/go_api_tutorial/internal/models"
	"example/go_api_tutorial/internal/repository/interfaces"
	"gorm.io/gorm"
)

// userRepository implements the UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetAll returns all users (excluding password)
func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Select("id", "username", "email", "role", "created_at", "updated_at").Find(&users).Error
	return users, err
}

// GetByID returns a user by ID (excluding password)
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Select("id", "username", "email", "role", "created_at", "updated_at").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername returns a user by username (including password for authentication)
func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail returns a user by email (including password for authentication)
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// UpdatePassword updates only the password of a user
func (r *userRepository) UpdatePassword(id uint, hashedPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("password", hashedPassword).Error
}

// Delete soft deletes a user
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// ExistsByUsername checks if a user exists by username
func (r *userRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// ExistsByEmail checks if a user exists by email
func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
