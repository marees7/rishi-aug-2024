package services

import "blogs/api/repositories"

type Services struct {
	AuthService  AuthServices
	UserService  UserServices
	AdminService AdminServices
}

func GetAuthService(repo *repositories.Repository) *Services {
	return &Services{
		AuthService: &authService{repo},
	}
}

func GetUserService(repo *repositories.Repository) *Services {
	return &Services{
		UserService: &userService{repo},
	}
}

func GetAdminService(repo *repositories.Repository) *Services {
	return &Services{
		AdminService: &adminService{repo},
	}
}
