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

type AdminHandler struct {
	services.AdminServices
}

// retrieve every users records
func (handler *AdminHandler) GetAllUsers(ctx echo.Context) error {
	var users []models.Users
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

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the get Users service
	status, err := handler.AdminServices.GetUsers(&users, role, limit, offset)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(status, helpers.ResponseJson{
		Message: "Users retreived successfully",
		Data:    users,
	})
}

// retrieve a single user record
func (handler *AdminHandler) GetSingleUser(ctx echo.Context) error {
	var user models.Users
	username := ctx.Param("username")

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the get User By ID service
	status, err := handler.AdminServices.GetUserByID(&user, username, role)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(status, helpers.ResponseJson{
		Message: "User retreived successfully",
		Data:    user,
	})
}

// create a new category
func (handler *AdminHandler) CreateCategory(ctx echo.Context) error {
	var category models.Categories

	if err := ctx.Bind(&category); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the create Category service
	status, err := handler.AdminServices.CreateCategory(&category, role)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(status, helpers.ResponseJson{
		Message: "Category created successfully",
		Data:    category,
	})
}

// update an existing category
func (handler *AdminHandler) UpdateCategory(ctx echo.Context) error {
	var category models.Categories
	id := ctx.Param("category_id")

	categoryid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&category); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the update category service
	if status, err := handler.AdminServices.UpdateCategory(&category, categoryid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Category updated successfully",
		Data: map[string]interface{}{
			"category_name": category.Category_name,
			"description":   category.Description,
			"updated_at":    category.UpdatedAt,
		},
	})
}

// delete an existing category
func (handler *AdminHandler) DeleteCategory(ctx echo.Context) error {
	var category models.Categories
	id := ctx.Param("category_id")

	categoryid, err := strconv.Atoi(id)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&category); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the delete category service
	if status, err := handler.AdminServices.DeleteCategory(&category, categoryid, role); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(status, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Category deleted successfully",
		Data: map[string]interface{}{
			"deleted_at": category.DeletedAt,
		},
	})
}
