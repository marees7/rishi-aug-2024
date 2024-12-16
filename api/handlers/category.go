package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Category services.CategoryServices
}

// retrieve every categories available
func (handler *CategoryHandler) Getcategories(ctx echo.Context) error {
	//call the retrieve category service
	categories, err := handler.Category.GetCategories()
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "retrieved categories successfully",
		Data:    categories,
	})
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

	if err := validation.CategoryValidation(&category); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if role == "admin" {
		//call the create Category service
		err := handler.Category.CreateCategory(&category)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
		}
	} else {
		return ctx.JSON(http.StatusUnauthorized, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Category created successfully",
		Data:    category,
	})
}

// update an existing category
func (handler *CategoryHandler) UpdateCategory(ctx echo.Context) error {
	var category models.Category
	id := (ctx.Param("category_id"))

	categoryid, err := uuid.Parse(id)
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

	if err := validation.CategoryValidation(&category); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if role == "admin" {
		//call the update category service
		if err := handler.Category.UpdateCategory(&category, categoryid); err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
		}
	} else {
		return ctx.JSON(http.StatusUnauthorized, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Category updated successfully",
		Data: map[string]interface{}{
			"category_name": category.Category_name,
			"description":   category.Description,
			"updated_at":    category.UpdatedAt,
		},
	})
}

// delete an existing category
func (handler *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	id := (ctx.Param("category_id"))

	categoryid, err := uuid.Parse(id)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	if role == "admin" {
		//call the delete category service
		category, err := handler.Category.DeleteCategory(categoryid, role)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message: "Category deleted successfully",
			Data: map[string]interface{}{
				"category_id": category.CategoryID,
				"deleted_at":  category.DeletedAt,
			},
		})
	} else {
		return ctx.JSON(http.StatusUnauthorized, dto.ResponseJson{
			Message: "Only admins are allowed",
		})

	}
}
