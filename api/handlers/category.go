package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/constants"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Category services.CategoryServices
}

// retrieve every categories available
func (handler *CategoryHandler) Getcategories(ctx echo.Context) error {
	//pagination
	pages := ctx.QueryParam("offset")
	page, err := strconv.Atoi(pages)
	if pages == "" {
		page = constants.DefaultOffset
	} else if err != nil {
		loggers.Warn.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, &dto.ResponseJson{
			Error: err.Error(),
		})
	}

	pageSize := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(pageSize)
	if pageSize == "" {
		limit = constants.DefaultLimit
	} else if err != nil {
		loggers.Warn.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, &dto.ResponseJson{
			Error: err.Error(),
		})
	}
	//call the retrieve category service
	categories, errs := handler.Category.GetCategories(limit, page)
	if errs != nil {
		loggers.Warn.Println(errs.Error)
		return ctx.JSON(errs.Status, dto.ResponseJson{
			Error: errs.Error,
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

	if err := validation.ValidateCategory(&category); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	roleCtx := ctx.Get("role").(string)

	if validation.CheckRole(roleCtx) {
		//call the create Category service
		err := handler.Category.CreateCategory(&category)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
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
	roleCtx := ctx.Get("role").(string)

	if validation.CheckRole(roleCtx) {
		//call the update category service
		if err := handler.Category.UpdateCategory(&category, categoryid); err != nil {
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
			"category_id": category.CategoryID,
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

	roleCtx := ctx.Get("role").(string)

	if validation.CheckRole(roleCtx) {
		//call the delete category service
		category, err := handler.Category.DeleteCategory(categoryid, roleCtx)
		if err != nil {
			loggers.Warn.Println(err.Error)
			return ctx.JSON(err.Status, dto.ResponseJson{
				Error: err.Error,
			})
		}

		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message: "Category deleted successfully",
			Data: map[string]interface{}{
				"category_id": category.CategoryID,
			},
		})
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})

	}
}
