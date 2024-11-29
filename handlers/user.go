package handlers

import (
	"blogs/helpers"
	"blogs/loggers"
	"blogs/models"
	"blogs/repositories"
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
	Getcategories(ctx echo.Context) error
}

type userHandler struct {
	services.UserServices
}

func (handler *userHandler) GetUsers(ctx echo.Context) error {
	var users []models.Users

	err := repositories.GetRepository().RetrieveUsers(&users)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	for _, user := range users {
		ctx.JSON(http.StatusOK, helpers.ResponseJson{
			Message: "Users retreived successfully",
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

	err := repositories.GetRepository().RetrieveSingleUser(&user, username)
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

	if err := repositories.GetRepository().CreatePost(&post, userid); err != nil {
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

	if err := repositories.GetRepository().UpdatePost(&post, userid, postid); err != nil {
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

func (handler *userHandler) Getcategories(ctx echo.Context) error {
	var categories []models.Categories

	err := repositories.GetRepository().RetrieveCategories(&categories)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	for _, category := range categories {
		ctx.JSON(http.StatusOK, helpers.ResponseJson{
			Message: "categories retreived successfully",
			Data: map[string]interface{}{
				"category id":   category.CategoryID,
				"category name": category.Category_name,
				"description":   category.Description,
			},
		})
	}
	return nil
}
