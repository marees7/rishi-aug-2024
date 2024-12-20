package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/dto"
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	services.CommentServices
}

// Create a new comment for a post
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
func (handler *CommentHandler) GetComments(ctx echo.Context) error {
	var postID uuid.UUID

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
	search := ctx.QueryParam("search")

	//pagination
	offsetStr := ctx.QueryParam("offset")
	limitStr := ctx.QueryParam("limit")

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
		PageSize:     limit,
		Page:         offset,
		TotalRecords: count})
}

// update an existing comment
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
func (handler *CommentHandler) DeleteComment(ctx echo.Context) error {
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

	roleCtx := ctx.Get("role").(string)

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
