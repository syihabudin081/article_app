package services

import "OrdentTest/models"

type ArticleService interface {
	CreateArticle(article *models.Article) error
	GetAllArticles() ([]*models.Article, error)
	GetArticleById(id string) (*models.Article, error)
	UpdateArticle(id string, article *models.Article) error
	DeleteArticle(id string) error
	GetArticlesByUserId(userId string) ([]*models.Article, error)
}
