package handlers

import (
	"blogs/helpers"
	"blogs/loggers"
	"blogs/models"
	"blogs/repositories"
	"blogs/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandlers interface {
	GetAllUsers(ctx echo.Context) error
	GetSingleUser(ctx echo.Context) error
}

type adminHandler struct {
	services.UserServices
}

func (handler *adminHandler) GetAllUsers(ctx echo.Context) error {
	var users []models.Users

	err := repositories.GetRepository().RetrieveUsers(&users)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Users retreived successfully",
		Data:    users,
	})

	return nil
}

func (handler *adminHandler) GetSingleUser(ctx echo.Context) error {
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
		Data:    user,
	})
}

func (handler *adminHandler) CreateCategory(ctx echo.Context) error {
	var category models.Categories

	if err := ctx.Bind(&category); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	err := repositories.GetRepository().CreateCategory(&category)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Category created successfully",
		Data:    category,
	})
}
