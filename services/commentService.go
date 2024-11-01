package services

import "OrdentTest/models"

type CommentService interface {
	CreateComment(comment *models.Comment) error
	GetAllComments() ([]*models.Comment, error)
	GetCommentById(id string) (*models.Comment, error)
	UpdateComment(id string, comment *models.Comment) error
	DeleteComment(id string) error
	GetCommentsByArticleId(articleId string) ([]*models.Comment, error)
	GetCommentsByUserId(userId string) ([]*models.Comment, error)
}
