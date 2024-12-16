package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type PostServices interface {
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post, postid uuid.UUID, role string) error
	DeletePost(userid uuid.UUID, postid uuid.UUID, role string) (*models.Post, error)
	GetPosts(startdate string, enddate string, postid uuid.UUID, title string) (*[]models.Post, error)
}

type postService struct {
	repositories.PostRepository
}

func InitPostService(post repositories.PostRepository) PostServices {
	return &postService{post}
}

// create a new post
func (repo postService) CreatePost(post *models.Post) error {
	if err := repo.PostRepository.CreatePost(post); err != nil {
		return err
	}
	return nil
}

// update a existing post
func (repo postService) UpdatePost(post *models.Post, postid uuid.UUID, role string) error {
	if err := repo.PostRepository.UpdatePost(post, postid, role); err != nil {
		return err
	}
	return nil
}

// delete a existing post
func (repo postService) DeletePost(userid uuid.UUID, postid uuid.UUID, role string) (*models.Post, error) {
	post, err := repo.PostRepository.DeletePost(userid, postid, role)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// retrieve every users posts using date or post id
func (repo postService) GetPosts(startdate string, enddate string, postid uuid.UUID, title string) (*[]models.Post, error) {
	//check if the given fields are valid
	post, err := repo.PostRepository.GetPosts(startdate, enddate, postid, title)
	if err != nil {
		return nil, err
	}
	return post, nil
}
