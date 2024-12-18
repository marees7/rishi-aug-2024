package routes

import (
	"blogs/api/handlers"
	"blogs/api/middlewares"
	"blogs/api/repositories"
	"blogs/api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryRoute(server *echo.Echo, db *gorm.DB) {
	//send the db connection to the repository package
	categoryRepository := repositories.InitCategoryRepository(db)

	//send the repo to the services package
	userService := services.InitCategoryService(categoryRepository)

	//Initialize the handler struct
	handler := &handlers.CategoryHandler{Category: userService}

	//group user routes
	users := server.Group("v1/users")
	users.Use(middlewares.ValidateToken)

	users.GET("/categories", handler.Getcategories)

	//group admin routes
	admin := server.Group("v1/admin")
	admin.Use(middlewares.ValidateToken)

	admin.POST("/categories", handler.CreateCategory)
	admin.PUT("/categories/:category_id", handler.UpdateCategory)
	admin.DELETE("/categories/:category_id", handler.DeleteCategory)
}
