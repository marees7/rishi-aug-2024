package handlers

import "blogs/services"

type Handlers struct {
	AdminHandlers
	AuthHandlers
	UserHandlers
}

func GetHandlerDB(service *services.Services) *Handlers {
	return &Handlers{
		AdminHandlers: &adminHandler{service.AdminService},
		AuthHandlers:  &authHandler{service.AuthService},
		UserHandlers:  &userHandler{service.UserService},
	}
}
