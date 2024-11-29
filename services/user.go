package services

import "blogs/repositories"

type UserServices interface {
	GetUsers()
	GetUserByID()
}

type userService struct {
	repository *repositories.Repository
}

func (repo userService) GetUsers() {

}

func (repo userService) GetUserByID() {

}
