package handlers

import (
	"github.com/marees7/rishi-aug-2024/api/services"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/pkg/loggers"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"net/http"

	"github.com/marees7/rishi-aug-2024/api/validation"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	services.AuthServices
}

// register an new user
//
// @Summary 	Register a new user
// @Description Creates and register a new user
// @Tags 		Auth
// @Accept 		json
// @produce 	json
// @param 		Signup  body models.User true "Enter your details"
// @success 	201 {object} dto.ResponseJson
// @failure		400 {object} dto.ResponseJson
// @failure		409 {object} dto.ResponseJson
// @failure		500 {object} dto.ResponseJson
// @router 		/signup [post]
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
//
// @Summary 	log in a new user
// @Description sign in a user and validate the token
// @Tags 		Auth
// @Accept 		json
// @produce 	json
// @Param   	Login      body dto.LoginRequest true "Enter your login details"
// @success 	200 {object} dto.ResponseJson
// @failure		400 {object} dto.ResponseJson
// @failure		404 {object} dto.ResponseJson
// @failure		500 {object} dto.ResponseJson
// @router 		/login [post]
func (handler *AuthHandler) Login(ctx echo.Context) error {
	var login dto.LoginRequest

	if err := ctx.Bind(&login); err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//check if the email field is empty
	if login.Email == "" {
		loggers.Warn.Println("email cannot be empty")
		return ctx.JSON(http.StatusBadRequest, dto.ResponseJson{
			Error: "email cannot be empty",
		})
	}

	//check if the password field is empty
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
	tokenStr, err := validation.GenerateToken(user)
	if err != nil {
		loggers.Warn.Println(err)
		return ctx.JSON(http.StatusUnauthorized, dto.ResponseJson{
			Error: err.Error(),
		})
	}

	//use the generated token to set a new cookie
	ctx.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    tokenStr,
		MaxAge:   1800,
		Secure:   false,
		HttpOnly: true,
	})

	return ctx.JSON(http.StatusOK, dto.ResponseJson{
		Message: "Logged in successfully",
	})
}
