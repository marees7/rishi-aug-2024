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

	//Initialize the handler struct
	handler := &handlers.AdminHandler{AdminServices: adminService}

	//group admin routes
	admin := server.Group("v1/admin/users")
	admin.Use(middlewares.ValidateToken)

	admin.GET("", handler.GetUsers)
	admin.GET("/:username", handler.GetUser)

	//group user routes
	user := server.Group("v1/users")
	user.Use(middlewares.ValidateToken)

	user.PUT("", handler.UpdateUser)
	user.DELETE("", handler.DeleteUser)
}
