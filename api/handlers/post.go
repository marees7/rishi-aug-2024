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

type PostHandler struct {
	services.PostServices
}

// create a new post
//
// @Summary 	Create post
// @Description Create a new post
// @ID 			Create-post
// @Tags 		Posts
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		Create_post  body models.Post true "Create a new post"
// @Success 	201 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/post [post]
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
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "post created successfully",
		Data:    post,
	})
}

// retrieve every users posts using date filter or specific id
//
// @Summary 	get posts
// @Description get all posts
// @ID 			get-posts
// @Tags 		Posts
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		limit  query string false "Enter the limit"
// @param 		offset  query string false "Enter the offset"
// @param 		startDate  query string false "Enter the start date"
// @param 		endDate  query string false "Enter the end date"
// @param 		title  query string false "Enter the title to search"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/post [get]
func (handler *PostHandler) GetPosts(ctx echo.Context) error {
	var postID uuid.UUID
	fromDate := ctx.QueryParam("start_date")
	toDate := ctx.QueryParam("end_date")
	offsetStr := ctx.QueryParam("offset")
	limitStr := ctx.QueryParam("limit")
	title := ctx.QueryParam("title")

	id := ctx.QueryParam("post_id")
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
		"fromDate": fromDate,
		"toDate":   toDate,
		"title":    title,
		"limit":    limit,
		"offset":   offset,
	}

	//call the retrieve post service
	posts, count, err := handler.PostServices.GetPosts(postID, keywords)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message:      "Posts retrieved successfully",
		Data:         posts,
		Limit:        limit,
		Offset:       offset,
		TotalRecords: count,
	})
}

// retrieve a specific post using id
//
// @Summary 	get post
// @Description get single posts
// @ID 			get-post
// @Tags 		Posts
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/post [get]
func (handler *PostHandler) GetPost(ctx echo.Context) error {
	postCtx := ctx.Param("post_id")
	postID, err := uuid.Parse(postCtx)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	post, errorResponse := handler.PostServices.GetPost(postID)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ErrorResponse{
			Error: errorResponse.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Post retrieved successfully",
		Data:    post,
	})
}

// update a existing post
//
// @Summary 	Update post
// @Description Update a specific post
// @ID 			update-post
// @Tags 		Posts
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		postID  path string true "Enter the post id"
// @param 		Update_post  body models.Post true "Update the post"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/post/{postID} [put]
func (handler *PostHandler) UpdatePost(ctx echo.Context) error {
	var post models.Post

	id := ctx.Param("post_id")
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
		Data:    map[string]interface{}{"post_id": postID},
	})
}

// Delete a existing post
//
// @Summary 	delete post
// @Description delete a specific post
// @ID 			delete-post
// @Tags 		Posts
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		postID  path string true "Enter the post id"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/post/{postID} [delete]
func (handler *PostHandler) DeletePost(ctx echo.Context) error {
	id := ctx.Param("post_id")
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
	errorResponse := handler.PostServices.DeletePost(userID, postID, roleCtx)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ResponseJson{
			Error: errorResponse.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Post deleted successfully",
		Data:    postID,
	})
}
