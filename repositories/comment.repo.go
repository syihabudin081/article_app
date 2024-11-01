package repositories

import (
	"OrdentTest/models"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	FindAll() ([]*models.Comment, error)
	FindById(id string) (*models.Comment, error)
	Update(id string, comment *models.Comment) error
	Delete(id string) error
	FindByArticleId(articleId string) ([]*models.Comment, error)
	FindByUserId(userId string) ([]*models.Comment, error)
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: db}
}

func (cr *CommentRepositoryImpl) Create(comment *models.Comment) error {
	return cr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}

		// Preload the related data after creating the comment
		return tx.Preload("User").
			Preload("Article").
			First(comment, comment.ID).Error
	})
}

func (cr *CommentRepositoryImpl) FindAll() ([]*models.Comment, error) {
	var comments []*models.Comment
	err := cr.DB.Preload("User").
		Preload("Article").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *CommentRepositoryImpl) FindById(id string) (*models.Comment, error) {
	var comment models.Comment
	err := cr.DB.Preload("User").
		Preload("Article").
		First(&comment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (cr *CommentRepositoryImpl) Update(id string, comment *models.Comment) error {
	return cr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Comment{}).Where("id = ?", id).Updates(comment).Error; err != nil {
			return err
		}

		// Preload updated data
		return tx.Preload("User").
			Preload("Article").
			First(comment, "id = ?", id).Error
	})
}

func (cr *CommentRepositoryImpl) Delete(id string) error {
	return cr.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&models.Comment{}).Error
	})
}

// FindByArticleId retrieves comments for a specific article ID
func (cr *CommentRepositoryImpl) FindByArticleId(articleId string) ([]*models.Comment, error) {
	// Validate UUID format
	if _, err := uuid.Parse(articleId); err != nil {
		return nil, errors.New("invalid article ID")
	}

	var comments []*models.Comment
	err := cr.DB.Preload("User").
		Preload("Article").
		Where("article_id = ?", articleId).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByUserId retrieves comments for a specific user ID
func (cr *CommentRepositoryImpl) FindByUserId(userId string) ([]*models.Comment, error) {
	// Validate UUID format
	log.Info("UserID : ", userId)
	if _, err := uuid.Parse(userId); err != nil {
		return nil, errors.New("invalid user ID")
	}

	var comments []*models.Comment
	err := cr.DB.Preload("User").
		Preload("Article").
		Where("user_id = ?", userId).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
