package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/dto"
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Category services.CategoryServices
}

// create a new category
func (handler *CategoryHandler) CreateCategory(ctx echo.Context) error {
	var category models.Category

	if err := ctx.Bind(&category); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.ValidateCategory(&category); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	roleCtx := ctx.Get("role").(string)

	if validation.ValidateRole(roleCtx) {
		//call the create Category service
		err := handler.Category.CreateCategory(&category)
		if err != nil {
			loggers.Warn.Println(err.Error)
			return ctx.JSON(err.Status, dto.ResponseJson{Error: err.Error})
		}
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "Category created successfully",
		Data:    category,
	})
}

// retrieve every categories available
func (handler *CategoryHandler) GetCategories(ctx echo.Context) error {
	offsetStr := ctx.QueryParam("offset")
	limitStr := ctx.QueryParam("limit")

	//pagination
	limit, offset, err := helpers.Pagination(limitStr, offsetStr)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//call the retrieve category service
	categories, errorResponse, count := handler.Category.GetCategories(limit, offset)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ResponseJson{
			Error: errorResponse.Error,
		})
	}
	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message:      "retrieved categories successfully",
		Data:         categories,
		PageSize:     limit,
		Page:         offset,
		TotalRecords: count,
	})
}

// update an existing category
func (handler *CategoryHandler) UpdateCategory(ctx echo.Context) error {
	var category models.Category
	id := (ctx.Param("category_id"))

	categoryID, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := ctx.Bind(&category); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	roleCtx := ctx.Get("role").(string)

	if validation.ValidateRole(roleCtx) {
		//call the update category service
		if err := handler.Category.UpdateCategory(&category, categoryID); err != nil {
			loggers.Warn.Println(err.Error)
			return ctx.JSON(err.Status, dto.ResponseJson{
				Error: err.Error,
			})
		}
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Category updated successfully",
		Data: map[string]interface{}{
			"category_id": categoryID,
		},
	})
}

// delete an existing category
func (handler *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	id := ctx.Param("category_id")

	categoryID, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	roleCtx := ctx.Get("role").(string)

	if validation.ValidateRole(roleCtx) {
		//call the delete category service
		err := handler.Category.DeleteCategory(categoryID)
		if err != nil {
			loggers.Warn.Println(err.Error)
			return ctx.JSON(err.Status, dto.ResponseJson{
				Error: err.Error,
			})
		}

		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message: "Category deleted successfully",
			Data: map[string]interface{}{
				"category_id": categoryID,
			},
		})
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})

	}
}
