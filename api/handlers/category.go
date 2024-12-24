package handlers

import (
	"fmt"
	"net/http"

	"github.com/marees7/rishi-aug-2024/api/services"
	"github.com/marees7/rishi-aug-2024/api/validation"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/common/helpers"
	"github.com/marees7/rishi-aug-2024/pkg/loggers"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Category services.CategoryServices
}

// create a new category
//
// @Summary 	Create categories
// @Description Create a new category
// @ID 			create-category
// @Tags 		Category
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		CreateDetails  body models.Category true "Enter the category details that need to be created"
// @Success 	201 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		409 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/admin/categories [post]
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

// Retrieve every categories available
//
// @Summary 	Get categories
// @Description Get all the available categories
// @ID 			get-category
// @Tags 		Category
// @Security 	JWT
// @Produce 	json
// @Param       limit query string false "Enter the limit"
// @Param       offset query string false "Enter the offset"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users/categories [get]
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
		Limit:        limit,
		Offset:       offset,
		TotalRecords: count,
	})
}

// update an existing category
//
// @Summary 	Update categories
// @Description Update a existing category
// @ID 			update-category
// @Tags 		Category
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		categoryID  path string true "Enter the category id"
// @param 		UpdateDetails  body models.Category true "Enter the changes need to be done"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/admin/categories/{categoryID} [put]
func (handler *CategoryHandler) UpdateCategory(ctx echo.Context) error {
	var category models.Category

	id := (ctx.Param("category_id"))
	fmt.Println(id)
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
//
// @Summary 	Delete categories
// @Description Delete a existing category
// @ID 			Delete-category
// @Tags 		Category
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @param 		categoryID  path string true "Enter the category id"
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/admin/categories/{categoryID} [delete]
func (handler *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	id := ctx.Param("category_id")
	fmt.Println("id:", id)
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
