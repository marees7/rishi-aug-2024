package main

import (
	"blogs/api/handlers"
	"blogs/api/repositories"
	"blogs/api/routes"
	"blogs/api/services"
	initializers "blogs/internals"
	"blogs/pkg/loggers"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnv()
	loggers.OpenLog()
	initializers.Connect()
	initializers.Migrate()
}

func main() {
	server := echo.New()
	fmt.Println("Starting the http server...")

	db := initializers.GetDB()
	authRepo := repositories.GetAuthRepository(db)
	userRepo := repositories.GetUserRepository(db)

	authService := services.GetAuthService(authRepo)
	adminService := services.GetAdminService(userRepo)
	userService := services.GetUserService(userRepo)

	authhandler := handlers.GetAuthHandler(authService)
	adminHandler := handlers.GetAdminHandler(adminService)
	userHandler := handlers.GetUserHandler(userService)

	routes.AuthRoute(server, authhandler)
	routes.UserRoute(server, userHandler)
	routes.AdminRoute(server, adminHandler)

	if err := server.Start(os.Getenv("HTTP_PORT")); err != nil {
		fmt.Println("Failed to start the server", err)
		loggers.ErrorLog.Fatalln("Failed to start the server", err)
	}
}
