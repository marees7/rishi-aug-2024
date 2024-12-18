package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/constants"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	services.CommentService
}

// Create a new comment for a post
func (handler *CommentHandler) CreateComment(ctx echo.Context) error {
	var comment models.Comment
	id := (ctx.Param("post_id"))

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
	if err := handler.CommentService.CreateComment(&comment); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "comment created successfully",
		Data:    comment,
	})
}

// update an existing comment
func (handler *CommentHandler) UpdateComment(ctx echo.Context) error {
	var comment models.Comment
	id := (ctx.Param("comment_id"))

	commentid, err := uuid.Parse(id)
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
	if err := handler.CommentService.UpdateComment(&comment, commentid); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "comment updated successfully",
		Data: map[string]interface{}{
			"coment_id": comment.CommentID,
		},
	})
}

// delete an existing comment
func (handler *CommentHandler) DeleteComment(ctx echo.Context) error {
	id := (ctx.Param("comment_id"))

	commentid, err := uuid.Parse(id)
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
	comment, errs := handler.CommentService.DeleteComment(userID, commentid, roleCtx)
	if errs != nil {
		loggers.Warn.Println(errs.Error)
		return ctx.JSON(errs.Status, dto.ResponseJson{
			Error: errs.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "comment deleted successfully",
		Data:    comment.CommentID,
	})
}

// retrieve every comments of the post
func (handler *CommentHandler) GetComments(ctx echo.Context) error {
	id := ctx.Param("post_id")

	postID, err := uuid.Parse(id)
	if id == "" {
		postID = uuid.Nil
	} else if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	search := ctx.QueryParam("search")

	//pagination
	pages := ctx.QueryParam("offset")
	page, err := strconv.Atoi(pages)
	if pages == "" {
		page = constants.DefaultOffset
	} else if err != nil {
		loggers.Warn.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, &dto.ResponseJson{
			Error: err.Error(),
		})
	}

	pageSize := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(pageSize)
	if pageSize == "" {
		limit = constants.DefaultLimit
	} else if err != nil {
		loggers.Warn.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, &dto.ResponseJson{
			Error: err.Error(),
		})
	}

	commentMap := map[string]interface{}{
		"search": search,
		"limit":  limit,
		"offset": page,
	}

	//call the retrieve comment service
	comment, errs := handler.CommentService.GetComments(postID, commentMap)
	if errs != nil {
		loggers.Warn.Println(errs.Error)
		return ctx.JSON(errs.Status, dto.ResponseJson{
			Error: errs.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Comments retrieved successfully",
		Data:    comment})
}
