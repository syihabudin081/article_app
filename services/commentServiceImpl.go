package services

import (
	"OrdentTest/models"
	"OrdentTest/repositories"
)

type CommentServiceImpl struct {
	repositories.CommentRepository
}

func NewCommentService(cr repositories.CommentRepository) *CommentServiceImpl {
	return &CommentServiceImpl{cr}
}

func (cs *CommentServiceImpl) CreateComment(comment *models.Comment) error {
	return cs.CommentRepository.Create(comment)
}

func (cs *CommentServiceImpl) GetAllComments() ([]*models.Comment, error) {
	return cs.CommentRepository.FindAll()
}

func (cs *CommentServiceImpl) GetCommentById(id string) (*models.Comment, error) {
	comments, err := cs.CommentRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cs *CommentServiceImpl) UpdateComment(id string, comment *models.Comment) error {
	return cs.CommentRepository.Update(id, comment)
}

func (cs *CommentServiceImpl) DeleteComment(id string) error {
	return cs.CommentRepository.Delete(id)
}

func (cs *CommentServiceImpl) GetCommentsByArticleId(articleId string) ([]*models.Comment, error) {
	return cs.CommentRepository.FindByArticleId(articleId)
}

func (cs *CommentServiceImpl) GetCommentsByUserId(userId string) ([]*models.Comment, error) {
	return cs.CommentRepository.FindByUserId(userId)
}
