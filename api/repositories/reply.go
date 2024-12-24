package repositories

import (
	"github.com/marees7/rishi-aug-2024/common/constants"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/pkg/models"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReplyRepository interface {
	CreateReply(reply *models.Reply) *dto.ErrorResponse
	UpdateReply(reply *models.Reply, replyID uuid.UUID) *dto.ErrorResponse
	DeleteReply(replyID uuid.UUID, userID uuid.UUID, role string) *dto.ErrorResponse
}

type replyRepository struct {
	*gorm.DB
}

func InitReplyRepository(db *gorm.DB) ReplyRepository {
	return &replyRepository{db}
}

// create a new comment
func (db *replyRepository) CreateReply(reply *models.Reply) *dto.ErrorResponse {
	//check if the comment exists
	data := db.Where("comment_id=?", reply.CommentID).First(&models.Comment{})
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: "comment does not exist"}
	} else if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}

	//create the comment
	data = db.Create(&reply)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}

	return nil
}

// updates the existing comment
func (db *replyRepository) UpdateReply(reply *models.Reply, replyID uuid.UUID) *dto.ErrorResponse {
	var replyData models.Reply

	//check if the record exists and if the user can access it
	data := db.Where("reply_id=?", replyID).First(&replyData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if replyData.UserID != reply.UserID {
		return &dto.ErrorResponse{Status: http.StatusForbidden, Error: "cannot update other users reply"}
	}

	//updates the reply if it is the user created it
	data = db.Where("reply_id=?", replyID).Updates(&reply)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}

	return nil
}

// deletes the existing comment
func (db *replyRepository) DeleteReply(replyID uuid.UUID, userID uuid.UUID, role string) *dto.ErrorResponse {
	var replyData models.Reply

	//check if the record exists and if the user can access it
	data := db.Where("reply_id=?", replyID).First(&replyData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if replyData.UserID != userID && role != constants.AdminRole {
		return &dto.ErrorResponse{Status: http.StatusForbidden, Error: "cannot delete other users reply"}
	}

	//deletes the record if the user created it or if it is the admin
	data = db.Where("reply_id=?", replyID).Delete(&replyData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}
	
	return nil
}
