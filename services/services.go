package services

import "blogs/repositories"

type Services struct {
	AuthService  AuthServices
	UserService  UserServices
	AdminService AdminServices
}

func GetService(repo *repositories.Repository) *Services {
	return &Services{
		AuthService:  &authService{repo},
		UserService:  &userService{repo},
		AdminService: &adminService{repo},
	}
}
