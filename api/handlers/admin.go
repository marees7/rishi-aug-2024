package handlers

import (
	"blogs/api/services"
	"blogs/api/validation"
	"blogs/common/constants"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	services.AdminServices
}

// retrieve every users records
func (handler *AdminHandler) GetUsers(ctx echo.Context) error {
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

	roleCtx := ctx.Get("role").(string)

	if validation.CheckRole(roleCtx) {
		//call the get Users service
		users, err := handler.AdminServices.GetUsers(limit, page)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message: "Users retreived successfully",
			Data:    users,
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

	if validation.CheckRole(roleCtx) {
		//call the get User By ID service
		users, err := handler.AdminServices.GetUser(username)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
				Error: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, dto.ResponseJson{
			Message: "Users retreived successfully",
			Data:    users,
		})
	} else {
		return ctx.JSON(http.StatusForbidden, dto.ResponseJson{
			Message: "Only admins are allowed",
		})
	}
}
