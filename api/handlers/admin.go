package handlers

import (
	"blogs/api/services"
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
	var limit, offset int

	limit_param := ctx.QueryParam("limit")
	if limit_param == "" {
		limit = constants.Default_Limit
	} else {
		convLimit, err := strconv.Atoi(limit_param)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
				Error: err.Error(),
			})
		}
		limit = convLimit
	}

	offset_param := ctx.QueryParam("offset")
	if offset_param == "" {
		offset = constants.Default_Offset
	} else {
		convOffset, err := strconv.Atoi(offset_param)
		if err != nil {
			loggers.Warn.Println(err)
			return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
				Error: err.Error(),
			})
		}
		offset = convOffset
	}

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the get Users service
	users, err := handler.AdminServices.GetUsers(role, limit, offset)
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
}

// retrieve a single user record
func (handler *AdminHandler) GetUser(ctx echo.Context) error {
	username := ctx.Param("username")

	getRole := (ctx.Get("role"))
	role := getRole.(string)

	//call the get User By ID service
	user, err := handler.AdminServices.GetUser(username, role)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "User retreived successfully",
		Data:    user,
	})
}
