package repositories

import (
	"OrdentTest/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Compare this snippet from services/commentServiceImpl.go:
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	return ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})

}

func (ur *userRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := ur.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(user *models.User) error {
	return ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func (ur *userRepository) DeleteUser(id string) error {
	return ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	// Use Preload to load the associated Role
	err := ur.DB.Preload("Role").First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
