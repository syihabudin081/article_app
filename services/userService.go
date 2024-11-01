package services

import "OrdentTest/models"

type UserService interface {
	// Compare this snippet from services/userServiceImpl.go:
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*models.User, error)
}
