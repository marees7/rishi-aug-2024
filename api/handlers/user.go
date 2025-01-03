package handlers

import (
	"net/http"

	"github.com/marees7/rishi-aug-2024/api/services"
	"github.com/marees7/rishi-aug-2024/api/validation"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/common/helpers"
	"github.com/marees7/rishi-aug-2024/pkg/loggers"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	services.AdminServices
}

// retrieve every users records
//
// @Summary 	get users
// @Description get every users records
// @ID 			get-users
// @Tags 		users
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/admin/users [get]
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
//
// @Summary 	get user
// @Description get a single user record
// @ID 			get-user
// @Tags 		users
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		403 {object} dto.ResponseJson
// @Failure		404 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/admin/users/{username} [get]
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
//
// @Summary 	update user
// @Description update a logged in user
// @ID 			update-user
// @Tags 		users
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users [put]
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
//
// @Summary 	delete user
// @Description delete a logged in user
// @ID 			delete-user
// @Tags 		users
// @Security 	JWT
// @Accept		json
// @Produce 	json
// @Success 	200 {object} dto.ResponseJson
// @Failure		400 {object} dto.ResponseJson
// @Failure		304 {object} dto.ResponseJson
// @Failure		500 {object} dto.ResponseJson
// @Router 		/v1/users [delete]
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
