package handlers

import (
	"net/http"

	"github.com/marees7/rishi-aug-2024/api/services"
	"github.com/marees7/rishi-aug-2024/api/validation"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/common/helpers"
	"github.com/marees7/rishi-aug-2024/pkg/loggers"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	services.CommentServices
}

// Create a new comment for a post
//
// @Summary 	Create comment
// @Description Create a new comment
// @ID 			Create-comment
// @Tags 		Comments
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		postID  path string true "Enter the post id"
// @param 		Create_comment  body models.Comment true "Enter the message you want add in the comment"
// @Success 	201 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/comment/{postID} [post]
func (handler *CommentHandler) CreateComment(ctx echo.Context) error {
	var comment models.Comment

	id := ctx.Param("post_id")
	postID, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	comment.PostID = postID
	if err := ctx.Bind(&comment); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.ValidateComment(&comment); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	userIDCtx := (ctx.Get("user_id")).(string)
	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	comment.UserID = userID
	//call the create comment service
	if err := handler.CommentServices.CreateComment(&comment); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "comment added successfully",
		Data:    comment,
	})
}

// retrieve every comments of the post
//
// @Summary 	Get comment
// @Description Get comments in a post
// @ID 			get-comment
// @Tags 		Comments
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		postID  path string true "Enter the post id"
// @param 		limit  query string false "Enter the limit"
// @param 		offset  query string false "Enter the offset"
// @param 		search  query string false "Enter a comment phrase you want to search"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/comment/{postID} [get]
func (handler *CommentHandler) GetComments(ctx echo.Context) error {
	var postID uuid.UUID
	offsetStr := ctx.QueryParam("offset")
	limitStr := ctx.QueryParam("limit")
	search := ctx.QueryParam("search")

	id := ctx.Param("post_id")
	if id == "" {
		postID = uuid.Nil
	} else {
		convPostID, err := uuid.Parse(id)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
				Error: err.Error(),
			})
		}
		postID = convPostID
	}

	//pagination
	limit, offset, err := helpers.Pagination(limitStr, offsetStr)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	keywords := map[string]interface{}{
		"search": search,
		"limit":  limit,
		"offset": offset,
	}

	//call the retrieve comment service
	comment, errorResponse, count := handler.CommentServices.GetComments(postID, keywords)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ResponseJson{
			Error: errorResponse.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message:      "Comments retrieved successfully",
		Data:         comment,
		Limit:        limit,
		Offset:       offset,
		TotalRecords: count})
}

// update an existing comment
//
// @Summary 	Update comment
// @Description Update comments in a post
// @ID 			update-comment
// @Tags 		Comments
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		commentID  path string true "Enter the comment id"
// @param 		Update_comment  body models.Comment true "Update the comment"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/comment/{commentID} [put]
func (handler *CommentHandler) UpdateComment(ctx echo.Context) error {
	var comment models.Comment

	id := (ctx.Param("comment_id"))
	commentID, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&comment); err != nil {
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

	comment.UserID = userID
	//call the update comment service
	if err := handler.CommentServices.UpdateComment(&comment, commentID); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "comment edited successfully",
		Data: map[string]interface{}{
			"comment_id": commentID,
		},
	})
}

// delete an existing comment
//
// @Summary 	delete comment
// @Description delete a specific comment
// @ID 			delete-comment
// @Tags 		Comments
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		commentID  path string true "Enter the comment id"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/comment/{commentID} [delete]
func (handler *CommentHandler) DeleteComment(ctx echo.Context) error {
	roleCtx := ctx.Get("role").(string)

	id := (ctx.Param("comment_id"))
	commentID, err := uuid.Parse(id)
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

	//call the delete comment service
	errorResponse := handler.CommentServices.DeleteComment(userID, commentID, roleCtx)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ResponseJson{
			Error: errorResponse.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "comment deleted successfully",
		Data:    commentID,
	})
}
