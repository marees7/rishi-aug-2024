package handlers

import (
	"net/http"

	"github.com/marees7/rishi-aug-2024/api/services"
	"github.com/marees7/rishi-aug-2024/api/validation"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/pkg/loggers"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReplyHandler struct {
	services.ReplyServices
}

// create a new reply
//
// @Summary 	Create reply
// @Description Create a new reply
// @ID 			Create-reply
// @Tags 		Replies
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		commentID  path string true "Enter the comment id"
// @param 		Create_reply  body models.Reply true "Enter the reply you want add in the comment"
// @Success 	201 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/reply/{commentID} [post]
func (handler *ReplyHandler) CreateReply(ctx echo.Context) error {
	var reply models.Reply

	if err := ctx.Bind(&reply); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.ValidateReply(&reply); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	userIDCtx := ctx.Get("user_id").(string)
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	commentIDCtx := ctx.Param("comment_id")
	commentID, err := uuid.Parse(commentIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	reply.CommentID = commentID
	reply.UserID = userID
	//call the create post service
	if err := handler.ReplyServices.CreateReply(&reply); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "reply added to the comment successfully",
		Data:    reply,
	})
}

// update a existing post
//
// @Summary 	Update reply
// @Description Update reply of a comment
// @ID 			update-reply
// @Tags 		Replies
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		replyID  path string true "Enter the reply id"
// @param 		Update_reply  body models.Reply true "Update the reply"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/reply/{replyID} [put]
func (handler *ReplyHandler) UpdateReply(ctx echo.Context) error {
	var reply models.Reply

	if err := ctx.Bind(&reply); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.ValidateReply(&reply); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	userIDCtx := ctx.Get("user_id").(string)
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	replyIDCtx := ctx.Param("reply_id")
	replyID, err := uuid.Parse(replyIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	reply.UserID = userID
	//call the create post service
	if err := handler.ReplyServices.UpdateReply(&reply, replyID); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "reply edited successfully",
		Data:    map[string]interface{}{"reply_id": replyID},
	})
}

// Delete a existing reply
//
// @Summary 	delete reply
// @Description delete a specific reply
// @ID 			delete-reply
// @Tags 		Replies
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		replyID  path string true "Enter the reply id"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/reply/{replyID} [delete]
func (handler *ReplyHandler) DeleteReply(ctx echo.Context) error {
	id := ctx.Param("reply_id")
	replyID, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	userIDCtx := ctx.Get("user_id").(string)
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	roleCtx := ctx.Get("role").(string)
	//call the delete reply service
	errorResponse := handler.ReplyServices.DeleteReply(replyID, userID, roleCtx)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ResponseJson{
			Error: errorResponse.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "reply deleted successfully",
		Data:    replyID,
	})
}
