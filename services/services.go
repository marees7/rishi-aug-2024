package services

import "blogs/repositories"

type authServices struct {
	authService AuthServices
	userService UserServices
}

func GetService() *authServices {
	repo := repositories.GetRepository()
	return &authServices{
		authService: &authService{repo},
		userService: &userService{repo},
	}
}
