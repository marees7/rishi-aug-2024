package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"
	"fmt"
	"net/http"
)

type UserServices interface {
	GetUsers(users *[]models.Users, limit, offset int) (int, error)
	GetUserByID(users *models.Users, username string) (int, error)
	CreatePost(post *models.Posts, userid int) (int, error)
	UpdatePost(post *models.Posts, userid int, postid int, role string) (int, error)
	DeletePost(post *models.Posts, userid int, postid int, role string) (int, error)
	RetrievePost(posts *[]models.Posts, startdate string, enddate string, postid int) (int, error)
	RetrieveCategories(categories *[]models.Categories) (int, error)
	CreateComment(comment *models.Comments, userid, postid int) (int, error)
	UpdateComment(comment *models.Comments, userid int, commentid int, role string) (int, error)
	DeleteComment(comment *models.Comments, userid int, commentid int, role string) (int, error)
	RetrieveComment(comments *[]models.Comments, postid int) (int, error)
}

type userService struct {
	*repositories.Repository
}

func (repo userService) GetUsers(users *[]models.Users, limit, offset int) (int, error) {
	if status, err := repo.User.RetrieveUsers(users, limit, offset); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo userService) GetUserByID(users *models.Users, username string) (int, error) {
	if status, err := repo.User.RetrieveSingleUser(users, username); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo userService) CreatePost(post *models.Posts, userid int) (int, error) {
	if status, err := repo.User.CreatePost(post, userid); err != nil {
		return status, err
	}
	return http.StatusCreated, nil
}

func (repo userService) UpdatePost(post *models.Posts, userid int, postid int, role string) (int, error) {
	if status, err := repo.User.UpdatePost(post, userid, postid, role); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo userService) DeletePost(post *models.Posts, userid int, postid int, role string) (int, error) {
	if status, err := repo.User.DeletePost(post, userid, postid, role); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo userService) RetrievePost(posts *[]models.Posts, startdate string, enddate string, postid int) (int, error) {
	if startdate == "" && enddate != "" {
		return http.StatusBadRequest, fmt.Errorf("start date is not specified")
	} else if startdate != "" && enddate == "" {
		return http.StatusBadRequest, fmt.Errorf("end date is not specified")
	} else if startdate == "" && enddate == "" && postid == 0 {
		return http.StatusBadRequest, fmt.Errorf("specify either start and end date or use post id")
	} else if status, err := repo.User.RetrievePost(posts, startdate, enddate, postid); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo userService) RetrieveCategories(categories *[]models.Categories) (int, error) {
	if status, err := repo.User.RetrieveCategories(categories); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo *userService) CreateComment(comment *models.Comments, userid, postid int) (int, error) {
	if status, err := repo.User.CreateComment(comment, userid, postid); err != nil {
		return status, err
	}
	return http.StatusCreated, nil
}

func (repo *userService) UpdateComment(comment *models.Comments, userid int, commentid int, role string) (int, error) {
	if status, err := repo.User.UpdateComment(comment, userid, commentid, role); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo *userService) DeleteComment(comment *models.Comments, userid int, commentid int, role string) (int, error) {
	if status, err := repo.User.DeleteComment(comment, userid, commentid, role); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (repo userService) RetrieveComment(comments *[]models.Comments, postid int) (int, error) {
	if status, err := repo.User.RetrieveComment(comments, postid); err != nil {
		return status, err
	}
	return http.StatusOK, nil
}
