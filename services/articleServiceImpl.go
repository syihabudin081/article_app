package services

import (
	"OrdentTest/models"
	"OrdentTest/repositories"
)

type ArticleServiceImpl struct {
	repositories.ArticleRepository
}

func NewArticleService(ar repositories.ArticleRepository) *ArticleServiceImpl {
	return &ArticleServiceImpl{ar}
}

func (as *ArticleServiceImpl) CreateArticle(article *models.Article) error {
	return as.ArticleRepository.Create(article)
}

func (as *ArticleServiceImpl) GetAllArticles() ([]*models.Article, error) {
	return as.ArticleRepository.FindAll()
}

func (as *ArticleServiceImpl) GetArticleById(id string) (*models.Article, error) {
	articles, err := as.ArticleRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleServiceImpl) UpdateArticle(id string, article *models.Article) error {
	return as.ArticleRepository.Update(id, article)
}

func (as *ArticleServiceImpl) DeleteArticle(id string) error {
	return as.ArticleRepository.Delete(id)
}

func (as *ArticleServiceImpl) GetArticlesByUserId(userId string) ([]*models.Article, error) {
	return as.ArticleRepository.FindByUserId(userId)
}
