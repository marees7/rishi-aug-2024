package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type PostServices interface {
	CreatePost(post *models.Post) *dto.ErrorResponse
	GetPosts(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Post, int64, error)
	GetPost(postID uuid.UUID) (*models.Post, *dto.ErrorResponse)
	UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse
	DeletePost(userID uuid.UUID, postID uuid.UUID, role string) *dto.ErrorResponse
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
func (repo postService) GetPosts(postID uuid.UUID, keywords map[string]interface{}) (*[]models.Post, int64, error) {
	return repo.PostRepository.GetPosts(postID, keywords)
}

// retrieve single user posts using title or post id
func (repo postService) GetPost(postID uuid.UUID) (*models.Post, *dto.ErrorResponse) {
	return repo.PostRepository.GetPost(postID)
}

// update a existing post
func (repo postService) UpdatePost(post *models.Post, postID uuid.UUID) *dto.ErrorResponse {
	return repo.PostRepository.UpdatePost(post, postID)
}

// delete a existing post
func (repo postService) DeletePost(userID uuid.UUID, postID uuid.UUID, role string) *dto.ErrorResponse {
	return repo.PostRepository.DeletePost(userID, postID, role)
}
