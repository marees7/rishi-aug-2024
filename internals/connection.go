package initializers

import (
	"blogs/pkg/loggers"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	fmt.Println("Connecting to the server...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the server", err)
		loggers.ErrorLog.Fatalln(err)
	}
	loggers.InfoLog.Println("Connected to the server")
	fmt.Println("Connected to the server successfully...")

	db = client
}

func GetDB() *gorm.DB {
	return db
}
