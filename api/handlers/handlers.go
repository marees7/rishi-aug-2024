package handlers

import "blogs/api/services"

type Handlers struct {
	AdminHandlers
	AuthHandlers
	UserHandlers
}

func GetAuthHandler(service *services.Services) *Handlers {
	return &Handlers{
		AuthHandlers: &authHandler{service.AuthService},
	}
}

func GetAdminHandler(service *services.Services) *Handlers {
	return &Handlers{
		AdminHandlers: &adminHandler{service.AdminService},
	}
}

func GetUserHandler(service *services.Services) *Handlers {
	return &Handlers{
		UserHandlers: &userHandler{service.UserService},
	}
}
