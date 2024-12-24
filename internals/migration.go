package internals

import (
	"github.com/marees7/rishi-aug-2024/pkg/loggers"
	"github.com/marees7/rishi-aug-2024/pkg/models"
)

//Migrate the model structs to the database
func (db connection) Migrate() {
	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Post{}, &models.Comment{}, &models.Reply{})
	if err != nil {
		loggers.Error.Fatalln(err)
	}
	
	loggers.Info.Println("Migrated tables successfully...")
}
