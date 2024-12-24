package routes

import (
	"github.com/marees7/rishi-aug-2024/api/handlers"
	"github.com/marees7/rishi-aug-2024/api/middlewares"
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ReplyRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	replyRepository := repositories.InitReplyRepository(db)

	//send the repo to the services package
	replyService := services.ReplyServices(replyRepository)

	//Initialize the handler struct
	handler := &handlers.ReplyHandler{ReplyServices: replyService}

	//group user routes
	users := server.Group("v1/users/reply")
	users.Use(middlewares.ValidateToken)

	users.POST("/:comment_id", handler.CreateReply)
	users.PUT("/:reply_id", handler.UpdateReply)
	users.DELETE("/:reply_id", handler.DeleteReply)
}
