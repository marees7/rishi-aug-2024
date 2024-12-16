package internals

import (
	"blogs/pkg/loggers"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	*gorm.DB
}

func Connect() *connection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		loggers.Error.Fatalln(err)
	}
	loggers.Info.Println("Connected to the server")

	return &connection{client}
}
