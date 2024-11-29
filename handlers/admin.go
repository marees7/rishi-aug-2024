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

type AdminHandlers interface {
	GetAllUsers(ctx echo.Context) error
	GetSingleUser(ctx echo.Context) error
	CreateCategory(ctx echo.Context) error
	UpdateCategory(ctx echo.Context) error
	DeleteCategory(ctx echo.Context) error
}

type adminHandler struct {
	services.AdminServices
}

func (handler *adminHandler) GetAllUsers(ctx echo.Context) error {
	var users []models.Users

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	err := handler.AdminServices.GetUsers(&users, role)
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

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	err := handler.AdminServices.GetUserByID(&user, username, role)
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

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	err := handler.AdminServices.CreateCategory(&category, role)
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

func (handler *adminHandler) UpdateCategory(ctx echo.Context) error {
	var category models.Categories
	id := ctx.Param("category_id")

	categoryid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid category id entered,check again",
			Error:   err.Error(),
		})
	}

	if err := ctx.Bind(&category); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if err := handler.AdminServices.UpdateCategory(&category, categoryid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Message: "Something went wrong",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Category updated successfully",
		Data:    category,
	})
}

func (handler *adminHandler) DeleteCategory(ctx echo.Context) error {
	var category models.Categories
	id := ctx.Param("category_id")

	categoryid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid category id entered,check again",
			Error:   err.Error(),
		})
	}

	if err := ctx.Bind(&category); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if err := handler.AdminServices.DeleteCategory(&category, categoryid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Message: "Something went wrong",
			Error:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Category deleted successfully",
		Data:    category.DeletedAt,
	})
}
