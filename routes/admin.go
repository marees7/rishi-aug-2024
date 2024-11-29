package routes

import (
	"blogs/handlers"
	"blogs/middlewares"

	"github.com/labstack/echo/v4"
)

func AdminRoute(server *echo.Echo) {
	handler := handlers.GetHandlerDB()

	admin := server.Group("/admin")
	admin.Use(middlewares.RequireAuth)

	admin.GET("/user", handler.GetAllUsers)
	admin.GET("/user/:username", handler.GetSingleUser)
	admin.POST("/category", handler.CreateCategory)
	admin.GET("/category", handler.Getcategories)
	admin.POST("/post", handler.CreatePost)
}
