package routes

import (
	"github.com/marees7/rishi-aug-2024/api/handlers"
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	authRepository := repositories.InitAuthRepository(db)

	//send the repo to the services package
	authService := services.InitAuthService(authRepository)

	//Initialize the handler struct
	handler := &handlers.AuthHandler{AuthServices: authService}

	server.POST("/signup", handler.Signup)
	server.POST("/login", handler.Login)
}
