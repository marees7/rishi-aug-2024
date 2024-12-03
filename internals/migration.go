package internals

import (
	"blogs/pkg/loggers"
	"blogs/pkg/models"
)

func Migrate() {
	db := GetDB()

	err := db.AutoMigrate(&models.Users{}, &models.Categories{}, &models.Posts{}, &models.Comments{})
	if err != nil {
		loggers.ErrorLog.Fatalln(err)
	}
	loggers.InfoLog.Println("Migrated tables successfully...")
}
