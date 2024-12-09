package internals

import (
	"blogs/pkg/loggers"
	"blogs/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Users{}, &models.Categories{}, &models.Posts{}, &models.Comments{})
	if err != nil {
		loggers.Error.Fatalln(err)
	}
	loggers.Info.Println("Migrated tables successfully...")
}
