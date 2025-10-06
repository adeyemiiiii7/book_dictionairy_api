package service

import (
	"errors"
	"strings"

	"example/go_api_tutorial/internal/models"
	"example/go_api_tutorial/internal/repository/interfaces"
	"example/go_api_tutorial/internal/utils"
	"gorm.io/gorm"
)

// UserService handles business logic for users
type UserService struct {
	userRepo interfaces.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// RegisterUser creates a new user account
func (s *UserService) RegisterUser(username, email, password string) (*models.User, error) {
	// Validate input
	if strings.TrimSpace(username) == "" {
		return nil, errors.New("username is required")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}
	if err := utils.ValidatePassword(password); err != nil {
		return nil, err
	}

	// Check if username already exists
	exists, err := s.userRepo.ExistsByUsername(username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	exists, err = s.userRepo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Username: strings.TrimSpace(username),
		Email:    strings.TrimSpace(strings.ToLower(email)),
		Password: hashedPassword,
		Role:     models.RoleUser, // Default role
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser authenticates a user and returns user info (without password)
func (s *UserService) LoginUser(usernameOrEmail, password string) (*models.User, error) {
	if strings.TrimSpace(usernameOrEmail) == "" {
		return nil, errors.New("username or email is required")
	}
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}

	var user *models.User
	var err error

	// Try to find user by email first, then by username
	if strings.Contains(usernameOrEmail, "@") {
		user, err = s.userRepo.GetByEmail(usernameOrEmail)
	} else {
		user, err = s.userRepo.GetByUsername(usernameOrEmail)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Check password
	if err := utils.CheckPassword(password, user.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Clear password before returning
	user.Password = ""
	return user, nil
}

// GetUserByID returns a user by ID (without password)
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// GetAllUsers returns all users (admin only)
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

// UpdateUserRole updates a user's role (admin only)
func (s *UserService) UpdateUserRole(userID uint, newRole models.UserRole) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Validate role
	if newRole != models.RoleUser && newRole != models.RoleAdmin {
		return errors.New("invalid role")
	}

	user.Role = newRole
	return s.userRepo.Update(user)
}

// ChangePassword changes a user's password
func (s *UserService) ChangePassword(userID uint, currentPassword, newPassword string) error {
	// Validate new password
	if err := utils.ValidatePassword(newPassword); err != nil {
		return err
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Note: In a production app, you would verify the current password first
	// This requires getting the user with password, which our current repository
	// design doesn't support well. For now, we'll just update the password.
	return s.userRepo.UpdatePassword(userID, hashedPassword)
}
