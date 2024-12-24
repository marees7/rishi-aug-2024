package repositories

import (
	"errors"
	"net/http"

	"github.com/marees7/rishi-aug-2024/common/constants"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) *dto.ErrorResponse
	GetPosts(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Post, int64, error)
	GetPost(postID uuid.UUID) (*models.Post, *dto.ErrorResponse)
	UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse
	DeletePost(userID uuid.UUID, postID uuid.UUID, role string) *dto.ErrorResponse
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
func (db *postRepository) GetPosts(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Post, int64, error) {
	var post []models.Post
	var count int64
	fromDate := keywords["fromDate"].(string)
	toDate := keywords["toDate"].(string)
	title := keywords["title"].(string)
	limit := keywords["limit"].(int)
	offset := keywords["offset"].(int)

	//check whether to use date filter or post id
	data := db.Model(post).Preload("Comments").Limit(limit).Offset(offset).Count(&count).Find(&post)
	if data.Error != nil {
		return nil, 0, data.Error
	}

	if fromDate != "" {
		db.Preload("Comments").Where("created_at >= ?", fromDate)
	} else if toDate != "" {
		db.Preload("Comments").Where("created_at <= ?", toDate)
	} else if title != "" {
		db.Preload("Comments").Where("title LIKE '%' || ? || '%' ", title)
	}

	return &post, count, nil
}

func (db *postRepository) GetPost(postID uuid.UUID) (*models.Post, *dto.ErrorResponse) {
	var post models.Post

	data := db.Where("post_id=?", postID).Preload("Comments").First(&post)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: "post not found"}
	} else if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
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
		return &dto.ErrorResponse{Status: http.StatusForbidden, Error: "cannot update other users post"}
	}

	//updates the record if the user created it or if it is the admin
	data = db.Where("post_id=?", postID).Updates(&post)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}

	return nil
}

// delete a existing post
func (db *postRepository) DeletePost(userID uuid.UUID, postID uuid.UUID, role string) *dto.ErrorResponse {
	var postData models.Post

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postID).First(&postData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if postData.UserID != userID && role != constants.AdminRole {
		return &dto.ErrorResponse{Status: http.StatusForbidden, Error: "cannot delete other users post"}
	}

	//deletes the record if the user created it or if it is the admin
	data = db.Where("post_id=?", postID).Delete(&postData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}

	return nil
}
