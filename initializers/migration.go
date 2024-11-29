package initializers

import (
	"blogs/loggers"
	"blogs/models"
	"fmt"
)

func Migrate() {
	db := GetDB()

	err := db.AutoMigrate(&models.Users{}, &models.Categories{}, &models.Posts{}, &models.Comments{})
	if err != nil {
		fmt.Println("Could not migrate tables", err)
		loggers.ErrorLog.Fatalln(err)
	}
	loggers.InfoLog.Println("Migrated tables successfully...")
	fmt.Println("Migrated tables successfully...")
}
