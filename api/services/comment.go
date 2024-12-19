package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(comment *models.Comment) *dto.ErrorResponse
	GetComments(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Comment, *dto.ErrorResponse)
	UpdateComment(comment *models.Comment, commentID uuid.UUID) *dto.ErrorResponse
	DeleteComment(userID uuid.UUID, commentID uuid.UUID, role string) (*models.Comment, *dto.ErrorResponse)
}

type commentService struct {
	repositories.CommentRepository
}

func InitCommentService(comment repositories.CommentRepository) CommentService {
	return &commentService{comment}
}

// create a new comment
func (repo *commentService) CreateComment(comment *models.Comment) *dto.ErrorResponse {
	return repo.CommentRepository.CreateComment(comment)
}

// retrieve comments using post id
func (repo *commentService) GetComments(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Comment, *dto.ErrorResponse) {
	keywords["offset"] = (keywords["offset"].(int) - 1) * keywords["limit"].(int)

	return repo.CommentRepository.GetComments(postID, keywords)
}

// update a existing comment
func (repo *commentService) UpdateComment(comment *models.Comment, commentID uuid.UUID) *dto.ErrorResponse {
	return repo.CommentRepository.UpdateComment(comment, commentID)
}

// delete the existing comment
func (repo *commentService) DeleteComment(userID uuid.UUID, commentID uuid.UUID, role string) (*models.Comment, *dto.ErrorResponse) {
	return repo.CommentRepository.DeleteComment(userID, commentID, role)
}
