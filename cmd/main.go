package main

import (
	"blogs/api/routes"
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
}

func main() {
	//create a instance of echo
	server := echo.New()

	//connect to the database and get the db
	db := internals.Connect()

	//migrate the model structs
	db.Migrate()

	//send the services to the handlers package
	routes.AuthRoute(server, db.DB)
	routes.CategoryRoute(server, db.DB)
	routes.AdminRoute(server, db.DB)
	routes.CommentRoute(server, db.DB)
	routes.PostRoute(server, db.DB)
	routes.ReplyRoute(server, db.DB)

	//start the server
	if err := server.Start(os.Getenv("HTTP_PORT")); err != nil {
		loggers.Error.Fatalln("Failed to start the server", err)
	}
}
