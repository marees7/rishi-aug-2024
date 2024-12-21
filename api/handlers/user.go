package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/dto"
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	services.AdminServices
}

// retrieve every users records
func (handler *AdminHandler) GetUsers(ctx echo.Context) error {
	offsetStr := ctx.QueryParam("offset")
	limitStr := ctx.QueryParam("limit")
	name := ctx.QueryParam("name")

	//pagination
	limit, offset, err := helpers.Pagination(limitStr, offsetStr)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	roleCtx := ctx.Get("role").(string)
	if validation.ValidateRole(roleCtx) {
		//call the get Users service
		users, count, err := handler.AdminServices.GetUsers(limit, offset, name)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message:      "Users retrieved successfully",
			Data:         users,
			Limit:        limit,
			Offset:       offset,
			TotalRecords: count,
		})
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}
}

// retrieve a single user record
func (handler *AdminHandler) GetUser(ctx echo.Context) error {
	username := ctx.Param("username")

	roleCtx := ctx.Get("role").(string)
	if validation.ValidateRole(roleCtx) {
		//call the get User By ID service
		users, err := handler.AdminServices.GetUser(username)
		if err != nil {
			loggers.Warn.Println(err.Error)
			return ctx.JSON(err.Status, dto.ResponseJson{Error: err.Error})
		}

		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message: "Users retrieved successfully",
			Data:    users,
		})
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}
}

// update a existing user
func (handler *AdminHandler) UpdateUser(ctx echo.Context) error {
	var user models.User

	email := ctx.Get("email").(string)
	if err := ctx.Bind(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	user.Email = email
	//check if the given info is valid
	if err := validation.ValidateUser(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//call the update user service
	if err := handler.AdminServices.UpdateUser(&user); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "user details updated successfully",
		Data:    email,
	})
}

// Delete a existing user
func (handler *AdminHandler) DeleteUser(ctx echo.Context) error {
	email := ctx.Get("email").(string)

	//call the update user service
	err := handler.AdminServices.DeleteUser(email)
	if err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{
			Error: err.Error,
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "user deleted successfully",
		Data:    email,
	})
}
