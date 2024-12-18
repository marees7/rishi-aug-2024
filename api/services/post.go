package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type PostServices interface {
	CreatePost(post *models.Post) *dto.ErrorResponse
	GetPosts(postID uuid.UUID, postMap map[string]interface{}) (*[]models.Post, error)
	GetPost(postID uuid.UUID) (*models.Post, error)
	UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse
	DeletePost(userID uuid.UUID, postID uuid.UUID, role string) (*models.Post, *dto.ErrorResponse)
}

type postService struct {
	repositories.PostRepository
}

func InitPostService(post repositories.PostRepository) PostServices {
	return &postService{post}
}

// create a new post
func (repo postService) CreatePost(post *models.Post) *dto.ErrorResponse {
	return repo.PostRepository.CreatePost(post)
}

// retrieve every users posts using date or post id
func (repo postService) GetPosts(postID uuid.UUID, postMap map[string]interface{}) (*[]models.Post, error) {
	postMap["offset"] = (postMap["offset"].(int) - 1) * postMap["limit"].(int)

	return repo.PostRepository.GetPosts(postID, postMap)
}

// retrieve single user posts using title or post id
func (repo postService) GetPost(postID uuid.UUID) (*models.Post, error) {
	return repo.PostRepository.GetPost(postID)
}

// update a existing post
func (repo postService) UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse {
	return repo.PostRepository.UpdatePost(post, postID)
}

// delete a existing post
func (repo postService) DeletePost(userID uuid.UUID, postID uuid.UUID, role string) (*models.Post, *dto.ErrorResponse) {
	return repo.PostRepository.DeletePost(userID, postID, role)
}
