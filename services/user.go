package services

import (
	"blogs/models"
	"blogs/repositories"
)

type UserServices interface {
	GetUsers(users *[]models.Users) error
	GetUserByID(users *models.Users, username string) error
	CreatePost(post *models.Posts, userid int) error
	UpdatePost(post *models.Posts, userid int, postid int, role string) error
	DeletePost(post *models.Posts, userid int, postid int, role string) error
	RetrievePostByDate(posts *[]models.Posts, startdate string, enddate string) error
	RetrievePostByID(posts *models.Posts, postid int) error
	RetrieveCategories(categories *[]models.Categories) error
	CreateComment(comment *models.Comments, userid int) error
	UpdateComment(comment *models.Comments, userid int, commentid int, role string) error
	DeleteComment(comment *models.Comments, userid int, commentid int, role string) error
	RetrieveComment(comments *[]models.Comments, postid int) error
}

type userService struct {
	*repositories.Repository
}

func (repo userService) GetUsers(users *[]models.Users) error {
	if err := repo.User.RetrieveUsers(users); err != nil {
		return err
	}
	return nil
}

func (repo userService) GetUserByID(users *models.Users, username string) error {
	if err := repo.User.RetrieveSingleUser(users, username); err != nil {
		return err
	}
	return nil
}

func (repo userService) CreatePost(post *models.Posts, userid int) error {
	if err := repo.User.CreatePost(post, userid); err != nil {
		return err
	}
	return nil
}

func (repo userService) UpdatePost(post *models.Posts, userid int, postid int, role string) error {
	if err := repo.User.UpdatePost(post, userid, postid, role); err != nil {
		return err
	}
	return nil
}

func (repo userService) DeletePost(post *models.Posts, userid int, postid int, role string) error {
	if err := repo.User.DeletePost(post, userid, postid, role); err != nil {
		return err
	}
	return nil
}

func (repo userService) RetrievePostByDate(posts *[]models.Posts, startdate string, enddate string) error {
	if err := repo.User.RetrievePostByDate(posts, startdate, enddate); err != nil {
		return err
	}
	return nil
}

func (repo userService) RetrievePostByID(posts *models.Posts, postid int) error {
	if err := repo.User.RetrievePostWithID(posts, postid); err != nil {
		return err
	}
	return nil
}

func (repo userService) RetrieveCategories(categories *[]models.Categories) error {
	if err := repo.User.RetrieveCategories(categories); err != nil {
		return err
	}
	return nil
}

func (repo *userService) CreateComment(comment *models.Comments, userid int) error {
	if err := repo.User.CreateComment(comment, userid); err != nil {
		return err
	}
	return nil
}

func (repo *userService) UpdateComment(comment *models.Comments, userid int, commentid int, role string) error {
	if err := repo.User.UpdateComment(comment, userid, commentid, role); err != nil {
		return err
	}
	return nil
}

func (repo *userService) DeleteComment(comment *models.Comments, userid int, commentid int, role string) error {
	if err := repo.User.DeleteComment(comment, userid, commentid, role); err != nil {
		return err
	}
	return nil
}

func (repo userService) RetrieveComment(comments *[]models.Comments, postid int) error {
	if err := repo.User.RetrieveComment(comments, postid); err != nil {
		return err
	}
	return nil
}
