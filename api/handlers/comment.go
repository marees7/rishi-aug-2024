package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"

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

	postid, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	comment.PostID = postid

	if err := ctx.Bind(&comment); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.CommentValidation(&comment); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	user := ctx.Get("user_id")
	user_str := user.(string)

	userid, err := uuid.Parse(user_str)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	comment.UserID = userid

	//call the create comment service
	if err := handler.CommentService.CreateComment(&comment); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
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

	user := ctx.Get("user_id")
	user_str := user.(string)

	userid, err := uuid.Parse(user_str)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	comment.UserID = userid

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the update comment service
	if err := handler.CommentService.UpdateComment(&comment, commentid, role); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "comment updated successfully",
		Data: map[string]interface{}{
			"content": comment.Content,
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

	user := ctx.Get("user_id")
	user_str := user.(string)

	userid, err := uuid.Parse(user_str)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the delete comment service
	comment, err := handler.CommentService.DeleteComment(userid, commentid, role)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "comment deleted successfully",
		Data:    comment.CommentID,
	})
}

// retrieve every comments of the post
func (handler *CommentHandler) GetComment(ctx echo.Context) error {
	id := ctx.QueryParam("post_id")

	postid, err := uuid.Parse(id)
	if id == "" {
		postid = uuid.Nil
	} else if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	content := ctx.QueryParam("find")

	//call the retrieve comment service
	comment, err := handler.CommentService.GetComment(postid, content)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Comments retrieved successfully",
		Data:    comment})
}
