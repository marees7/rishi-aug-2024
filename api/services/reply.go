package services

import (
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/google/uuid"
)

type ReplyServices interface {
	CreateReply(reply *models.Reply) *dto.ErrorResponse
	UpdateReply(reply *models.Reply, replyID uuid.UUID) *dto.ErrorResponse
	DeleteReply(replyID uuid.UUID, userID uuid.UUID, role string) *dto.ErrorResponse
}

type replyService struct {
	repositories.ReplyRepository
}

func InitReplyService(reply repositories.ReplyRepository) ReplyServices {
	return &replyService{reply}
}

// create a new comment
func (repo *replyService) CreateReply(reply *models.Reply) *dto.ErrorResponse {
	return repo.ReplyRepository.CreateReply(reply)
}

// update a existing comment
func (repo *replyService) UpdateComment(reply *models.Reply, replyID uuid.UUID) *dto.ErrorResponse {
	return repo.ReplyRepository.UpdateReply(reply, replyID)
}

// delete the existing comment
func (repo *replyService) DeleteComment(replyID uuid.UUID, userID uuid.UUID, role string) *dto.ErrorResponse {
	return repo.ReplyRepository.DeleteReply(replyID, userID, role)
}
