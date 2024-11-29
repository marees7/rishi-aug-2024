package handlers

import (
	"blogs/helpers"
	"blogs/loggers"
	"blogs/models"
	"blogs/services"
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
	GetPostsByID(ctx echo.Context) error
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

	err := handler.UserServices.GetUsers(&users)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Message: "Something went wrong",
			Error:   err.Error(),
		})
	}

	for _, user := range users {
		ctx.JSON(http.StatusOK, helpers.ResponseJson{
			Data: map[string]interface{}{
				"username": user.Username,
				"posts":    user.Posts,
				"comments": user.Comments,
			},
		})
	}
	return nil
}

func (handler *userHandler) GetUserWithUsername(ctx echo.Context) error {
	var user models.Users

	username := ctx.Param("username")

	err := handler.UserServices.GetUserByID(&user, username)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "User retreived successfully",
		Data: map[string]interface{}{
			"username": user.Username,
			"posts":    user.Posts,
			"comments": user.Comments,
		},
	})
}

func (handler *userHandler) CreatePost(ctx echo.Context) error {
	var post models.Posts

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	if err := handler.UserServices.CreatePost(&post, userid); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "User retreived successfully",
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
			Message: "Invalid post id entered,check again",
			Error:   err.Error(),
		})
	}

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if err := handler.UserServices.UpdatePost(&post, userid, postid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Post updated successfully",
		Data:    post,
	})
}

func (handler *userHandler) DeletePost(ctx echo.Context) error {
	var post models.Posts
	id := ctx.Param("post_id")

	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid post id entered,check again",
			Error:   err.Error(),
		})
	}

	if err := ctx.Bind(&post); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if err := handler.UserServices.DeletePost(&post, userid, postid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Post deleted successfully",
	})
}

func (handler *userHandler) GetPosts(ctx echo.Context) error {
	var posts []models.Posts

	startDate := ctx.Param("start_date")
	endDate := ctx.Param("end_date")

	err := handler.UserServices.RetrievePostByDate(&posts, startDate, endDate)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Posts retrieved successfully",
		Data:    posts,
	})
}

func (handler *userHandler) GetPostsByID(ctx echo.Context) error {
	var posts models.Posts

	id := ctx.Param("post_id")
	postid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	err = handler.UserServices.RetrievePostByID(&posts, postid)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Posts retrieved successfully",
		Data:    posts,
	})
}

func (handler *userHandler) Getcategories(ctx echo.Context) error {
	var categories []models.Categories

	err := handler.UserServices.RetrieveCategories(&categories)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	for _, category := range categories {
		ctx.JSON(http.StatusOK, helpers.ResponseJson{
			Data: map[string]interface{}{
				"category id":   category.CategoryID,
				"category name": category.Category_name,
				"description":   category.Description,
			},
		})
	}
	return nil
}

func (handler *userHandler) CreateComment(ctx echo.Context) error {
	var comment models.Comments

	if err := ctx.Bind(&comment); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	if err := handler.UserServices.CreateComment(&comment, userid); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
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
			Message: "Invalid comment id entered,check again",
			Error:   err.Error(),
		})
	}

	if err := ctx.Bind(&comment); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if err := handler.UserServices.UpdateComment(&comment, userid, commentid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "comment updated successfully",
		Data:    comment,
	})
}

func (handler *userHandler) DeleteComment(ctx echo.Context) error {
	var comment models.Comments
	id := ctx.Param("comment_id")

	commentid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid comment id entered,check again",
			Error:   err.Error(),
		})
	}

	if err := ctx.Bind(&comment); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}
	user := ctx.Get("user_id")
	userid := int(user.(float64))

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if err := handler.UserServices.DeleteComment(&comment, userid, commentid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
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
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	err = handler.UserServices.RetrieveComment(&comments, postid)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
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
