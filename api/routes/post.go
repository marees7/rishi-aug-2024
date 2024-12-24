package routes

import (
	"github.com/marees7/rishi-aug-2024/api/handlers"
	"github.com/marees7/rishi-aug-2024/api/middlewares"
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PostRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	postRepository := repositories.InitPostRepository(db)

	//send the repo to the services package
	postService := services.InitPostService(postRepository)

	//Initialize the handler struct
	handler := &handlers.PostHandler{PostServices: postService}

	//group user routes
	users := server.Group("v1/users/post")
	users.Use(middlewares.ValidateToken)

	users.POST("", handler.CreatePost)
	users.GET("", handler.GetPosts)
	users.GET("/:post_id", handler.GetPost)
	users.PUT("/:post_id", handler.UpdatePost)
	users.DELETE("/:post_id", handler.DeletePost)
}
