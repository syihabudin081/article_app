package repositories

import (
	"OrdentTest/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll() ([]*models.Article, error)
	FindById(id string) (*models.Article, error)
	Create(article *models.Article) error
	Update(id string, article *models.Article) error
	Delete(id string) error
	FindByUserId(userId string) ([]*models.Article, error)
}

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{DB: db}
}

func (ar *articleRepository) Create(article *models.Article) error {
	return ar.DB.Transaction(func(tx *gorm.DB) error {
		// Create new article
		if err := tx.Create(article).Error; err != nil {
			return err
		}

		// Preload related data after creation
		return tx.Preload("Author").
			Preload("Author.Role").
			First(article, article.ID).Error
	})
}

func (ar *articleRepository) FindAll() ([]*models.Article, error) {
	var articles []*models.Article
	err := ar.DB.Preload("Author").
		Preload("Author.Role").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (ar *articleRepository) FindById(id string) (*models.Article, error) {
	var article models.Article
	err := ar.DB.Preload("Author").
		Preload("Author.Role").
		First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (ar *articleRepository) Update(id string, article *models.Article) error {
	return ar.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).Where("id = ?", id).Updates(article).Error; err != nil {
			return err
		}

		// Preload updated data
		return tx.Preload("Author").
			Preload("Author.Role").
			First(article, "id = ?", id).Error
	})
}

func (ar *articleRepository) Delete(id string) error {
	return ar.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&models.Article{}).Error
	})
}

func (ar *articleRepository) FindByUserId(userId string) ([]*models.Article, error) {
	var articles []*models.Article
	err := ar.DB.Preload("Author").
		Preload("Author.Role").
		Where("author_id = ?", userId).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
