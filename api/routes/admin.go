package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"

	"github.com/labstack/echo/v4"
)

func AdminRoute(server *echo.Echo, handler *handlers.Handlers) {
	admin := server.Group("/admin")
	admin.Use(middlewares.RequireAuth)

	admin.GET("/user", handler.GetAllUsers)
	admin.GET("/user/:username", handler.GetSingleUser)
	admin.POST("/categories", handler.CreateCategory)
	admin.PUT("/categories/:category_id", handler.UpdateCategory)
	admin.DELETE("/categories/:category_id", handler.DeleteCategory)
}
