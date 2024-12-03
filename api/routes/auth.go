package routes

import (
	"blogs/api/handlers"
	"blogs/api/services"

	"github.com/labstack/echo/v4"
)

func AuthRoute(server *echo.Echo, service *services.Services) {
	handler := &handlers.AuthHandler{AuthServices: service.AuthService}

	server.POST("/signup", handler.Signup)
	server.POST("/login", handler.Login)
}
