package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"
	"blogs/api/repositories"
	"blogs/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommentRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	commentRepository := repositories.InitCommentRepository(db)

	//send the repo to the services package
	commentService := services.InitCommentService(commentRepository)

	//Initialize the handler struct
	handler := &handlers.CommentHandler{CommentService: commentService}

	//group user routes
	users := server.Group("v1/users")
	users.Use(middlewares.ValidateToken)

	users.POST("/comments/:post_id", handler.CreateComment)
	users.GET("/comments/:post_id", handler.GetComments)
	users.PUT("/comments/:comment_id", handler.UpdateComment)
	users.DELETE("/comments/:comment_id", handler.DeleteComment)
}
