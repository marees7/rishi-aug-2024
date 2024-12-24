package routes

import (
	"github.com/marees7/rishi-aug-2024/api/handlers"
	"github.com/marees7/rishi-aug-2024/api/middlewares"
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/api/services"

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
	users := server.Group("v1/users/categories")
	users.Use(middlewares.ValidateToken)

	users.GET("", handler.GetCategories)

	//group admin routes
	admin := server.Group("v1/admin/categories")
	admin.Use(middlewares.ValidateToken)

	admin.POST("", handler.CreateCategory)
	admin.PUT("/:category_id", handler.UpdateCategory)
	admin.DELETE("/:category_id", handler.DeleteCategory)
}
