package repositories

import (
	"blogs/api/validation"
	"blogs/common/constants"
	"blogs/common/dto"
	"blogs/pkg/models"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) *dto.ErrorResponse
	GetPosts(postID uuid.UUID, postMap map[string]interface{}) (*[]models.Post, error)
	GetPost(postID uuid.UUID) (*models.Post, error)
	UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse
	DeletePost(userID uuid.UUID, postID uuid.UUID, role string) (*models.Post, *dto.ErrorResponse)
}

type postRepository struct {
	*gorm.DB
}

func InitPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

// create a new post
func (db *postRepository) CreatePost(post *models.Post) *dto.ErrorResponse {
	//check if the category exists
	data := db.Where("category_id=?", post.CategoryID).First(&models.Category{})
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: "category not found"}
	} else if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}

	//creates a new post
	data = db.Create(&post)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}
	return nil
}

// retrieve every users posts using either date or post id
func (db *postRepository) GetPosts(postID uuid.UUID, postMap map[string]interface{}) (*[]models.Post, error) {
	var post []models.Post
	startDate := postMap["startDate"].(string)
	endDate := postMap["endDate"].(string)
	title := postMap["title"].(string)
	limit := postMap["limit"].(int)
	offset := postMap["offset"].(int)

	//check whether to use date filter or post id
	if startDate != "" && endDate != "" && postID != uuid.Nil {
		data := db.Where("created_at BETWEEN ? AND ? AND post_id= ?", startDate, endDate, postID).Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	} else if startDate != "" && endDate != "" {
		data := db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	} else if startDate != "" && endDate == "" {
		data := db.Where("created_at > ?", startDate).Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	} else if startDate == "" && endDate != "" {
		data := db.Where("created_at < ?", endDate).Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	} else if postID != uuid.Nil {
		data := db.Where("post_id=?", postID).Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	} else if title != "" {
		data := db.Where("title LIKE '%' || ? || '%' ", title).Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	} else {
		data := db.Preload("Comments").Limit(limit).Offset(offset).Find(&post)
		if data.Error != nil {
			return nil, data.Error
		}
	}
	return &post, nil
}

func (db *postRepository) GetPost(postID uuid.UUID) (*models.Post, error) {
	var post models.Post

	data := db.Where("post_id=?", postID).Preload("Comments").First(&post)
	if data.Error != nil {
		if data.Error != nil {
			return nil, data.Error
		}
	}
	return &post, nil
}

// update a existing post
func (db *postRepository) UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse {
	var postData models.Post

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postID).First(&postData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if postData.UserID != post.UserID {
		return &dto.ErrorResponse{Status: http.StatusUnauthorized, Error: "cannot update other users post"}

	}

	//updates the record if the user created it or if it is the admin
	if post.UserID == postData.UserID {
		data := db.Where("post_id=?", postID).Updates(&post)
		if data.Error != nil {
			return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
		} else if data.RowsAffected == 0 {
			return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
		}
	}
	post.PostID = postData.PostID

	return nil
}

// delete a existing post
func (db *postRepository) DeletePost(userID uuid.UUID, postID uuid.UUID, role string) (*models.Post, *dto.ErrorResponse) {
	var postData models.Post

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postID).First(&postData)
	if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if postData.UserID != userID && role != constants.AdminRole {
		return nil, &dto.ErrorResponse{Status: http.StatusUnauthorized, Error: "cannot delete other users post"}
	}

	//deletes the record if the user created it or if it is the admin
	if postData.UserID == userID || validation.CheckRole(role) {
		data := db.Where("post_id=?", postID).Delete(&postData)
		if data.Error != nil {
			return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
		} else if data.RowsAffected == 0 {
			return nil, &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
		}
	}
	return &postData, nil
}
