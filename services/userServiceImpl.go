package services

import (
	"OrdentTest/models"
	"OrdentTest/repositories"
	"fmt"
)

// userService struct
type userService struct {
	repo repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo}
}

// GetAllUsers retrieves all users with pagination
func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAllUsers()
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserById(idStr string) (*models.User, error) {
	// Convert string ID to int

	user, err := s.repo.GetUserById(idStr)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return user, nil
}

// CreateUser creates a new user in the database
func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

// GetUserByEmail fetches a user by email
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found by email: %v", err)
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(idStr string) error {

	return s.repo.DeleteUser(idStr)
}
