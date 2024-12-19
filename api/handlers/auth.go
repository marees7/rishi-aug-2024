package handlers

import (
	"blogs/api/services"
	"blogs/common/dto"
	"blogs/pkg/loggers"
	"blogs/pkg/models"

	"blogs/api/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	services.AuthServices
}

// register an new user
func (handler *AuthHandler) Signup(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//check if the given info is valid
	if err := validation.ValidateUser(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//call the signup service
	if err := handler.AuthServices.Signup(&user); err != nil {
		loggers.Warn.Println(err.Error)
		return ctx.JSON(err.Status, dto.ResponseJson{Error: err.Error})
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseJson{
		Message: "User created successfully",
	})
}

// validate and sign-in a user
func (handler *AuthHandler) Login(ctx echo.Context) error {
	var login dto.LoginRequest

	if err := ctx.Bind(&login); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	if login.Email == "" {
		loggers.Warn.Println("email cannot be empty")
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: "email cannot be empty",
		})
	}

	if login.Password == "" {
		loggers.Warn.Println("password cannot be empty")
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: "password cannot be empty",
		})
	}

	//call the login service
	user, errorResponse := handler.AuthServices.Login(&login)
	if errorResponse != nil {
		loggers.Warn.Println(errorResponse.Error)
		return ctx.JSON(errorResponse.Status, dto.ResponseJson{
			Error: errorResponse.Error,
		})
	}

	//generate a new token
	tokenstr, err := validation.GenerateToken(user)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusUnauthorized, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//use the generated token to set a new cookie
	ctx.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    tokenstr,
		MaxAge:   1800,
		Secure:   false,
		HttpOnly: true,
	})

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Logged in successfully",
	})
}
