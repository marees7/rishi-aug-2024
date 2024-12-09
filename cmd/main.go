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
}

func main() {
	//create a instance of echo
	server := echo.New()

	//get the db connection
	db := internals.GetDB()
	//migrate the model structs
	internals.Migrate(db)
	//send the db connection to the repository package
	authRepository := repositories.GetAuthRepository(db)
	userRepository := repositories.GetUserRepository(db)
	//send the repo to the services package
	authService := services.GetAuthService(authRepository)
	adminService := services.GetAdminService(userRepository)
	userService := services.GetUserService(userRepository)
	//send the services to the handlers package
	routes.AuthRoute(server, authService)
	routes.UserRoute(server, userService)
	routes.AdminRoute(server, adminService)

	//start the server
	if err := server.Start(os.Getenv("HTTP_PORT")); err != nil {
		loggers.Error.Fatalln("Failed to start the server", err)
	}
}
