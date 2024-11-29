package main

import (
	"blogs/handlers"
	"blogs/initializers"
	"blogs/loggers"
	"blogs/repositories"
	"blogs/routes"
	"blogs/services"
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
	repo := repositories.GetRepository(db)
	service := services.GetService(repo)
	handler := handlers.GetHandlerDB(service)

	routes.AuthRoute(server, handler)
	routes.UserRoute(server, handler)
	routes.AdminRoute(server, handler)

	if err := server.Start(os.Getenv("http_port")); err != nil {
		fmt.Println("Failed to start the server", err)
		loggers.ErrorLog.Fatalln("Failed to start the server", err)
	}
}
