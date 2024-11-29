package routes

import (
	"blogs/handlers"
	"blogs/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(server *echo.Echo) {
	handler := handlers.GetHandlerDB()
	users := server.Group("/user")
	users.Use(middlewares.RequireAuth)

	users.GET("/user", handler.GetUsers)
	users.GET("/user/:username", handler.GetUserWithUsername)
	users.GET("/category", handler.Getcategories)
	users.POST("/post", handler.CreatePost)
	users.PUT("/post/:post_id", handler.UpdatePost)
}
