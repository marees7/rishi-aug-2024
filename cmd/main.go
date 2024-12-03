package main

import (
	"blogs/api/repositories"
	"blogs/api/routes"
	"blogs/api/services"
	"blogs/internals"
	"blogs/pkg/loggers"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	//load the env file
	internals.LoadEnv()
	//load the logger
	loggers.OpenLog()
	//connect to the database
	internals.Connect()
	//migrate the model structs
	internals.Migrate()
}

func main() {
	//create a instance of echo
	server := echo.New()

	//get the db connection
	db := internals.GetDB()
	//send the db connection to the repository package
	authRepo := repositories.GetAuthRepository(db)
	userRepo := repositories.GetUserRepository(db)
	//send the repo to the services package
	authService := services.GetAuthService(authRepo)
	adminService := services.GetAdminService(userRepo)
	userService := services.GetUserService(userRepo)
	//send the services to the handlers package
	routes.AuthRoute(server, authService)
	routes.UserRoute(server, userService)
	routes.AdminRoute(server, adminService)

	//start the server
	if err := server.Start(os.Getenv("HTTP_PORT")); err != nil {
		loggers.ErrorLog.Fatalln("Failed to start the server", err)
	}
}
