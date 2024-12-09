package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"
	"blogs/api/services"

	"github.com/labstack/echo/v4"
)

func UserRoute(server *echo.Echo, service *services.Services) {
	handler := &handlers.UserHandler{UserServices: service.UserService}

	users := server.Group("v1/users")
	users.Use(middlewares.RequireAuth)

	users.GET("/", handler.GetUsers)
	users.GET("/:username", handler.GetUser)
	users.GET("/categories", handler.Getcategories)
	users.POST("/posts", handler.CreatePost)
	users.GET("/posts", handler.GetPosts)
	users.PUT("/posts/:post_id", handler.UpdatePost)
	users.DELETE("/posts/:post_id", handler.DeletePost)
	users.POST("/comments/:post_id", handler.CreateComment)
	users.PUT("/comments/:comment_id", handler.UpdateComment)
	users.DELETE("/comments/:comment_id", handler.DeleteComment)

}
