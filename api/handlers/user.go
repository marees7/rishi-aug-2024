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

type UserHandlers interface {
	GetUsers(ctx echo.Context) error
	GetUserWithUsername(ctx echo.Context) error
	CreatePost(ctx echo.Context) error
	UpdatePost(ctx echo.Context) error
	DeletePost(ctx echo.Context) error
	GetPosts(ctx echo.Context) error
	Getcategories(ctx echo.Context) error
	CreateComment(ctx echo.Context) error
	UpdateComment(ctx echo.Context) error
	DeleteComment(ctx echo.Context) error
}

type userHandler struct {
	services.UserServices
}

func (handler *userHandler) GetUsers(ctx echo.Context) error {
	var users []models.Users
	var limit, offset int

	param := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(param)
	if err != nil {
		if param == "" {
			limit = 100
		} else {
			loggers.WarningLog.Println(err)
			return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Error: err.Error(),
			})
		}
	}

	param = ctx.QueryParam("offset")
	offset, err = strconv.Atoi(param)
	if err != nil {
		if param == "" {
			offset = 0
		} else {
			loggers.WarningLog.Println(err)
			return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
				Error: err.Error(),
			})
		}
	}

	status, err := handler.UserServices.GetUsers(&users, limit, offset)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	for _, user := range users {
		ctx.JSON(http.StatusOK, helpers.ResponseJson{
			Data: map[string]interface{}{
				"posts":    user.Posts,
				"comments": user.Comments,
				"username": user.Username,
				"email":    user.Email,
				"user_id":  user.UserID,
			},
		})
	}
	return nil
}

func (handler *userHandler) GetUserWithUsername(ctx echo.Context) error {
	var user models.Users

	username := ctx.Param("username")

	status, err := handler.UserServices.GetUserByID(&user, username)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "User retreived successfully",
		Data: map[string]interface{}{
			"posts":    user.Posts,
			"comments": user.Comments,
			"username": user.Username,
			"email":    user.Email,
			"user_id":  user.UserID,
		},
	})
}

func (handler *userHandler) CreatePost(ctx echo.Context) error {
	var post models.Posts

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

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

func (handler *userHandler) UpdatePost(ctx echo.Context) error {
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

func (handler *userHandler) DeletePost(ctx echo.Context) error {
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

func (handler *userHandler) GetPosts(ctx echo.Context) error {
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

func (handler *userHandler) Getcategories(ctx echo.Context) error {
	var categories []models.Categories

	status, err := handler.UserServices.RetrieveCategories(&categories)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (handler *userHandler) CreateComment(ctx echo.Context) error {
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

func (handler *userHandler) UpdateComment(ctx echo.Context) error {
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

func (handler *userHandler) DeleteComment(ctx echo.Context) error {
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

func (handler *userHandler) GetComments(ctx echo.Context) error {
	var comments []models.Comments

	id := ctx.Param("post_id")

	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

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
