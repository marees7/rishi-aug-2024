package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"
	"blogs/api/repositories"
	"blogs/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	userRepository := repositories.InitUserRepository(db)
	//send the repo to the services package
	adminService := services.InitAdminService(userRepository)

	handler := &handlers.AdminHandler{AdminServices: adminService}

	//group admin routes
	admin := server.Group("v1/admin")
	admin.Use(middlewares.TokenValidation)

	admin.GET("/users", handler.GetUsers)
	admin.GET("/users/:username", handler.GetUser)
}
