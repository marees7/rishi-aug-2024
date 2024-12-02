package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(server *echo.Echo, handler *handlers.Handlers) {
	users := server.Group("/users")
	users.Use(middlewares.RequireAuth)

	users.GET("/user", handler.GetUsers)
	users.GET("/user/:username", handler.GetUserWithUsername)
	users.GET("/categories", handler.Getcategories)
	users.POST("/posts", handler.CreatePost)
	users.GET("/posts", handler.GetPosts)
	users.PUT("/posts/:post_id", handler.UpdatePost)
	users.DELETE("/posts/:post_id", handler.DeletePost)
	users.POST("/comments/:post_id", handler.CreateComment)
	users.PUT("/comments/:comment_id", handler.UpdateComment)
	users.DELETE("/comments/:comment_id", handler.DeleteComment)

}
