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

	if err := validation.ValidatePost(&post); err != nil {
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
	post.UserID = userID

	//call the create post service
	if err := handler.PostServices.CreatePost(&post); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
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

	postID, err := uuid.Parse(id)
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

	userIDCtx := ctx.Get("user_id").(string)

	userID, err := uuid.Parse(userIDCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	post.UserID = userID

	//call the update post service
	if err := handler.PostServices.UpdatePost(&post, postID); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Post updated successfully",
		Data:    post.PostID,
	})
}

// Delete a existing post
func (handler *PostHandler) DeletePost(ctx echo.Context) error {
	id := (ctx.Param("post_id"))

	postID, err := uuid.Parse(id)
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

	//call the delete post service
	post, errs := handler.PostServices.DeletePost(userID, postID, roleCtx)
	if errs != nil {
		loggers.Warn.Println(errs.Error)
		return ctx.JSON(errs.Status, dto.ResponseJson{
			Error: errs.Error,
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

	postID, err := uuid.Parse(id)
	if id == "" {
		postID = uuid.Nil
	} else if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	title := ctx.QueryParam("title")

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
	postMap := map[string]interface{}{
		"startDate": startDate,
		"endDate":   endDate,
		"title":     title,
		"limit":     limit,
		"offset":    page,
	}

	//call the retrieve post service
	posts, err := handler.PostServices.GetPosts(postID, postMap)
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

func (handler *PostHandler) GetPost(ctx echo.Context) error {
	postCtx := ctx.Param("post_id")

	postID, err := uuid.Parse(postCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	post, err := handler.PostServices.GetPost(postID)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Post retrieved successfully",
		Data:    post,
	})
}
