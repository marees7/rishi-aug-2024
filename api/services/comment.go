package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(comment *models.Comment) error
	UpdateComment(comment *models.Comment, commentid uuid.UUID, role string) error
	DeleteComment(userid uuid.UUID, commentid uuid.UUID, role string) (*models.Comment, error)
	GetComment(postid uuid.UUID, content string) (*[]models.Comment, error)
}

type commentService struct {
	repositories.CommentRepository
}

func InitCommentService(comment repositories.CommentRepository) CommentService {
	return &commentService{comment}
}

// create a new comment
func (repo *commentService) CreateComment(comment *models.Comment) error {
	if err := repo.CommentRepository.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

// update a existing comment
func (repo *commentService) UpdateComment(comment *models.Comment, commentid uuid.UUID, role string) error {
	if err := repo.CommentRepository.UpdateComment(comment, commentid, role); err != nil {
		return err
	}
	return nil
}

// delete the existing comment
func (repo *commentService) DeleteComment(userid uuid.UUID, commentid uuid.UUID, role string) (*models.Comment, error) {
	comment, err := repo.CommentRepository.DeleteComment(userid, commentid, role)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// retrieve comments using post id
func (repo *commentService) GetComment(postid uuid.UUID, content string) (*[]models.Comment, error) {
	comment, err := repo.CommentRepository.GetComment(postid, content)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
