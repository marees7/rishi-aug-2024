package routes

import (
	"blogs/handlers"
	"blogs/middlewares"

	"github.com/labstack/echo/v4"
)

func AdminRoute(server *echo.Echo, handler *handlers.Handlers) {
	admin := server.Group("/admin")
	admin.Use(middlewares.RequireAuth)

	admin.GET("/user", handler.GetAllUsers)
	admin.GET("/user/:username", handler.GetSingleUser)
	admin.POST("/category", handler.CreateCategory)
	admin.PUT("/category/:category_id", handler.UpdateCategory)
	admin.GET("/category", handler.Getcategories)
	admin.DELETE("/category/:category_id", handler.DeleteCategory)
}
