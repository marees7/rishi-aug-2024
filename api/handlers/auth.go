package handlers

import (
	"blogs/api/services"
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/pkg/models"

	"blogs/api/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandlers interface {
	Signup(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type authHandler struct {
	services.AuthServices
}

func (handler *authHandler) Signup(ctx echo.Context) error {
	var user models.Users

	if err := ctx.Bind(&user); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := validation.Validation(&user); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if err := handler.AuthServices.Signup(&user); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "User created successfully",
		Data:    user.Email,
	})
}

func (handler *authHandler) Login(ctx echo.Context) error {
	var login helpers.LoginRequest

	if err := ctx.Bind(&login); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	user, err := handler.AuthServices.Login(&login)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	tokenstr, err := validation.GenerateToken(user)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    tokenstr,
		MaxAge:   1800,
		Secure:   false,
		HttpOnly: true,
	})

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Logged in successfully",
	})
}
