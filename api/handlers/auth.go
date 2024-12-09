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

type AuthHandler struct {
	services.AuthServices
}

// register an new user
func (handler *AuthHandler) Signup(ctx echo.Context) error {
	var user models.Users

	if err := ctx.Bind(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	//check if the given info is valid
	if err := validation.Validation(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	//call the signup service
	if err := handler.AuthServices.Signup(&user); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "User created successfully",
		Data:    user.Email,
	})
}

// validate and sign-in a user
func (handler *AuthHandler) Login(ctx echo.Context) error {
	var login helpers.LoginRequest

	if err := ctx.Bind(&login); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	if login.Email == "" {
		loggers.Warn.Println("email cannot be empty")
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: "email cannot be empty",
		})
	}

	if login.Password == "" {
		loggers.Warn.Println("password cannot be empty")
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: "password cannot be empty",
		})
	}

	//call the login service
	user, err := handler.AuthServices.Login(&login)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
			Error: err.Error(),
		})
	}

	//generate a new token
	tokenstr, err := validation.GenerateToken(user)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helpers.ResponseJson{
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

	return ctx.JSON(http.StatusOK, helpers.ResponseJson{
		Message: "Logged in successfully",
	})
}
