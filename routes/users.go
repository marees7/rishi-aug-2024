package routes

import (
	"blogs/handlers"
	"blogs/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(server *echo.Echo, handler *handlers.Handlers) {
	users := server.Group("/user")
	users.Use(middlewares.RequireAuth)

	users.GET("/user", handler.GetUsers)
	users.GET("/user/:username", handler.GetUserWithUsername)
	users.GET("/category", handler.Getcategories)
	users.GET("/post/:start_date/:end_date", handler.GetPosts)
	users.GET("/post/:post_id", handler.GetPostsByID)
	users.POST("/post", handler.CreatePost)
	users.PUT("/post/:post_id", handler.UpdatePost)
	users.DELETE("/post/:post_id", handler.DeletePost)
	users.POST("/comment", handler.CreateComment)
	users.PUT("/comment/:comment_id", handler.UpdateComment)
	users.DELETE("/comment/:comment_id", handler.DeleteComment)

}
