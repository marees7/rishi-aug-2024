package routes

import (
	"github.com/marees7/rishi-aug-2024/api/handlers"
	"github.com/marees7/rishi-aug-2024/api/middlewares"
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommentRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	commentRepository := repositories.InitCommentRepository(db)

	//send the repo to the services package
	commentService := services.InitCommentService(commentRepository)

	//Initialize the handler struct
	handler := &handlers.CommentHandler{CommentServices: commentService}

	//group user routes
	users := server.Group("v1/users/comment")
	users.Use(middlewares.ValidateToken)

	users.POST("/:post_id", handler.CreateComment)
	users.GET("/:post_id", handler.GetComments)
	users.PUT("/:comment_id", handler.UpdateComment)
	users.DELETE("/:comment_id", handler.DeleteComment)
}
