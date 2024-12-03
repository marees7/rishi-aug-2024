package handlers

import (
	"blogs/api/services"
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	services.UserServices
}

// retrieve every users public fields
func (handler *UserHandler) GetUsers(ctx echo.Context) error {
	var users []models.Users
	var view []helpers.UserView
	var limit, offset int

	param := ctx.QueryParam("limit")
	if param == "" {
		limit = 100
	} else {
		convLimit, err := strconv.Atoi(param)
		if err != nil {
			loggers.WarningLog.Println(err)
			return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Error: err.Error(),
			})
		}
		limit = convLimit
	}

	param = ctx.QueryParam("offset")
	if param == "" {
		offset = 0
	} else {
		convOffset, err := strconv.Atoi(param)
		if err != nil {
			loggers.WarningLog.Println(err)
			return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Error: err.Error(),
			})
		}
		offset = convOffset
	}

	//call the get Users service
	status, err := handler.UserServices.GetUsers(&users, limit, offset)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	//used another struct to show only public fields
	for _, user := range users {
		publicFields := helpers.UserView{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Comments: user.Comments,
			Posts:    user.Posts,
		}
		view = append(view, publicFields)
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "retrieved records successfully",
		Data:    view,
	})
}

// retrieve single user public fields using username
func (handler *UserHandler) GetUserWithUsername(ctx echo.Context) error {
	var user models.Users

	username := ctx.Param("username")

	//call the get User by id service
	status, err := handler.UserServices.GetUserByID(&user, username)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "User retreived successfully",
		Data: &helpers.UserView{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Comments: user.Comments,
			Posts:    user.Posts,
		},
	})
}

// create a new post
func (handler *UserHandler) CreatePost(ctx echo.Context) error {
	var post models.Posts

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	//call the create post service
	if status, err := handler.UserServices.CreatePost(&post, userid); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, helpers.ResponseJson{
		Message: "post created successfully",
		Data:    post,
	})
}

// update a existing post
func (handler *UserHandler) UpdatePost(ctx echo.Context) error {
	var post models.Posts
	id := ctx.Param("post_id")

	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the update post service
	if status, err := handler.UserServices.UpdatePost(&post, userid, postid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Post updated successfully",
		Data: map[string]interface{}{
			"title":       post.Title,
			"content":     post.Content,
			"description": post.Description,
		},
	})
}

// Delete a existing post
func (handler *UserHandler) DeletePost(ctx echo.Context) error {
	var post models.Posts
	id := ctx.Param("post_id")

	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the delete post service
	if status, err := handler.UserServices.DeletePost(&post, userid, postid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Post deleted successfully",
	})
}

// retrieve every users posts using date filter or specific id
func (handler *UserHandler) GetPosts(ctx echo.Context) error {
	var posts []models.Posts

	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")

	id := ctx.QueryParam("post_id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		if id == "" {
			postid = 0
		} else {
			loggers.WarningLog.Println(err)
			return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Error: err.Error(),
			})
		}
	}

	//call the retrieve post service
	status, err := handler.UserServices.RetrievePost(&posts, startDate, endDate, postid)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Posts retrieved successfully",
		Data:    posts,
	})
}

// retrieve every categories available
func (handler *UserHandler) Getcategories(ctx echo.Context) error {
	var categories []models.Categories

	//call the retrieve category service
	status, err := handler.UserServices.RetrieveCategories(&categories)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, categories)
}

// Create a new comment for a post
func (handler *UserHandler) CreateComment(ctx echo.Context) error {
	var comment models.Comments
	id := ctx.Param("post_id")

	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&comment); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	//call the create comment service
	if status, err := handler.UserServices.CreateComment(&comment, userid, postid); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "comment created successfully",
		Data:    comment,
	})
}

// update an existing comment
func (handler *UserHandler) UpdateComment(ctx echo.Context) error {
	var comment models.Comments
	id := ctx.Param("comment_id")

	commentid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&comment); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the update comment service
	if status, err := handler.UserServices.UpdateComment(&comment, userid, commentid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "comment updated successfully",
		Data: map[string]interface{}{
			"content": comment.Content,
		},
	})
}

// delete an existing comment
func (handler *UserHandler) DeleteComment(ctx echo.Context) error {
	var comment models.Comments
	id := ctx.Param("comment_id")

	commentid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&comment); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the delete comment service
	if status, err := handler.UserServices.DeleteComment(&comment, userid, commentid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "comment deleted successfully",
	})
}

// retrieve every comments of the post
func (handler *UserHandler) GetComments(ctx echo.Context) error {
	var comments []models.Comments

	id := ctx.Param("post_id")

	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	//call the retrieve comment service
	status, err := handler.UserServices.RetrieveComment(&comments, postid)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	for _, comment := range comments {
		ctx.JSON(http.StatusOK, helpers.ResponseJson{
			Data: map[string]interface{}{
				"comment id":  comment.CommentID,
				"content":     comment.Content,
				"date posted": comment.CreatedAt,
			},
		})
	}
	return nil
}
