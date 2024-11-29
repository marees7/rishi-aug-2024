package main

import (
	"blogs/initializers"
	"blogs/loggers"
	"blogs/routes"
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

	routes.AuthRoute(server)
	routes.UserRoute(server)
	routes.AdminRoute(server)

	if err := server.Start(os.Getenv("http_port")); err != nil {
		fmt.Println("Failed to start the server", err)
		loggers.ErrorLog.Fatalln("Failed to start the server", err)
	}
}
