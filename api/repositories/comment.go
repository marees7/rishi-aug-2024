package repositories

import (
	"blogs/common/constants"
	"blogs/common/dto"
	"blogs/pkg/models"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) *dto.ErrorResponse
	GetComments(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Comment, *dto.ErrorResponse, int64)
	UpdateComment(comment *models.Comment, commentID uuid.UUID) *dto.ErrorResponse
	DeleteComment(userID uuid.UUID, commentID uuid.UUID, role string) *dto.ErrorResponse
}

type commentRepository struct {
	*gorm.DB
}

func InitCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

// create a new comment
func (db *commentRepository) CreateComment(comment *models.Comment) *dto.ErrorResponse {
	//check if the post exists
	data := db.Where("post_id=?", comment.PostID).First(&models.Post{})
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: "post does not exist"}
	} else if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}

	//create the comment
	data = db.Create(&comment)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}

	return nil
}

// retrieve comments using post id
func (db *commentRepository) GetComments(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Comment, *dto.ErrorResponse, int64) {
	var comment []models.Comment
	var count int64
	search := keywords["search"].(string)
	limit := keywords["limit"].(int)
	offset := keywords["offset"].(int)

	//check if the post exists
	data := db.Where("post_id=?", postID).First(&models.Post{})
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: "post not found"}, 0
	} else if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}, 0
	}

	//retrieve the comments
	data = db.Model(comment).Preload("Replies").Where("post_id=?", postID).Limit(limit).Offset(offset).Count(&count).Find(&comment)
	if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}, 0
	}
	if search != "" {
		//get the comments using content
		db.Where("content LIKE '%' || ? || '%'", search)
	}

	return &comment, nil, count
}

// updates the existing comment
func (db *commentRepository) UpdateComment(comment *models.Comment, commentID uuid.UUID) *dto.ErrorResponse {
	var commentData models.Comment

	//check if the record exists and if the user can access it
	data := db.Where("comment_id=?", commentID).First(&commentData)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: "post not found"}
	} else if commentData.UserID != comment.UserID {
		return &dto.ErrorResponse{Status: http.StatusForbidden, Error: "cannot update other users comment"}
	}

	//updates the comment if it is the user created it
	data = db.Preload("Reply").Where("comment_id=?", commentID).Updates(&comment)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}

	return nil
}

// deletes the existing comment
func (db *commentRepository) DeleteComment(userID uuid.UUID, commentID uuid.UUID, role string) *dto.ErrorResponse {
	var commentData models.Comment

	//check if the record exists and if the user can access it
	data := db.Where("comment_id=?", commentID).First(&commentData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if commentData.UserID != userID && role != constants.AdminRole {
		return &dto.ErrorResponse{Status: http.StatusForbidden, Error: "cannot delete other users comment"}
	}

	//deletes the record if the user created it or if it is the admin
	data = db.Where("comment_id=?", commentID).Delete(&commentData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}
	
	return nil
}
