package handlers

import (
	"blogs/helpers"
	"blogs/loggers"
	"blogs/models"
	"blogs/services"

	"blogs/validation"
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
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	if err := validation.Validation(&user); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusUnauthorized, helpers.ResponseJson{
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	if err := handler.AuthServices.Signup(&user); err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Message: "Something went wrong",
			Error:   err.Error(),
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
			Message: "Invalid data entered,check again",
			Error:   err.Error(),
		})
	}

	user, err := handler.AuthServices.Login(&login)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusBadGateway, helpers.ResponseJson{
			Message: "Invalid email or password entered,check again",
			Error:   err.Error(),
		})
	}

	tokenstr, err := validation.GenerateToken(user)
	if err != nil {
		loggers.WarningLog.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Message: "Could not generate token",
			Error:   err.Error(),
		})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    tokenstr,
		MaxAge:   800,
		Secure:   false,
		HttpOnly: true,
	})

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Logged in successfully",
	})
}
