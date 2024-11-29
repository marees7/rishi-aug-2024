package routes

import (
	"blogs/handlers"

	"github.com/labstack/echo/v4"
)

func AuthRoute(server *echo.Echo) {
	handler := handlers.GetHandlerDB()

	server.POST("/signup", handler.Signup)
	server.POST("/login", handler.Login)
}
