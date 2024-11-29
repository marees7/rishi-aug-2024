package services

import "blogs/repositories"

type AuthServices interface {
	Signup()
	Login()
}

type authService struct {
	repository *repositories.Repository
}

func (repo authService) Signup() {

}

func (repo authService) Login() {

}
