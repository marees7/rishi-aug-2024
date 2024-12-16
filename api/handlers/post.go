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

type PostHandler struct {
	services.PostServices
}

// create a new post
func (handler *PostHandler) CreatePost(ctx echo.Context) error {
	var post models.Post

	if err := ctx.Bind(&post); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.PostsValidation(&post); err != nil {
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
	post.UserID = userid

	//call the create post service
	if err := handler.PostServices.CreatePost(&post); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "post created successfully",
		Data:    post,
	})
}

// update a existing post
func (handler *PostHandler) UpdatePost(ctx echo.Context) error {
	var post models.Post
	id := (ctx.Param("post_id"))

	postid, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&post); err != nil {
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
	post.UserID = userid

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the update post service
	if err := handler.PostServices.UpdatePost(&post, postid, role); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Post updated successfully",
		Data: map[string]interface{}{
			"title":       post.Title,
			"content":     post.Content,
			"description": post.Description,
		},
	})
}

// Delete a existing post
func (handler *PostHandler) DeletePost(ctx echo.Context) error {
	id := (ctx.Param("post_id"))

	postid, err := uuid.Parse(id)
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

	//call the delete post service
	post, err := handler.PostServices.DeletePost(userid, postid, role)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Post deleted successfully",
		Data:    post.PostID,
	})
}

// retrieve every users posts using date filter or specific id
func (handler *PostHandler) GetPosts(ctx echo.Context) error {
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")

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

	title := ctx.QueryParam("title")

	if startDate == "" && endDate != "" {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: "End date is not specified",
		})
	} else if startDate != "" && endDate == "" {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: "Start date is not specified",
		})
	}

	//call the retrieve post service
	posts, err := handler.PostServices.GetPosts(startDate, endDate, postid, title)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Posts retrieved successfully",
		Data:    posts,
	})
}
