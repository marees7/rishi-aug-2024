package main

import (
	"os"

	"github.com/marees7/rishi-aug-2024/api/routes"
	"github.com/marees7/rishi-aug-2024/internals"
	"github.com/marees7/rishi-aug-2024/pkg/loggers"

	"github.com/labstack/echo/v4"
	_ "github.com/marees7/rishi-aug-2024/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	//load the env file
	internals.LoadEnv()

	//load the logger
	loggers.OpenLog()
}

// @title Blog posts API
// @version 0.0.1
// @description This is a blog post api where users can create posts and add comments for it

// @contact.name admin
// @contact.url https://github.com/marees7/rishi-aug-2024.git
// @contact.email rsi28c@gmail.com

// @host localhost:5030
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

	server.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	//start the server
	if err := server.Start(os.Getenv("HTTP_PORT")); err != nil {
		loggers.Error.Fatalln("Failed to start the server", err)
	}
}
