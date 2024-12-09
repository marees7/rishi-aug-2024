package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"
	"blogs/api/services"

	"github.com/labstack/echo/v4"
)

func AdminRoute(server *echo.Echo, service *services.Services) {
	handler := &handlers.AdminHandler{AdminServices: service.AdminService}

	admin := server.Group("v1/admin")
	admin.Use(middlewares.RequireAuth)

	admin.GET("/users", handler.GetUsers)
	admin.GET("/users/:username", handler.GetUser)
	admin.POST("/categories", handler.CreateCategory)
	admin.PUT("/categories/:category_id", handler.UpdateCategory)
	admin.DELETE("/categories/:category_id", handler.DeleteCategory)
}
