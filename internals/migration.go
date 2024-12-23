package internals

import (
	"blogs/pkg/loggers"
	"blogs/pkg/models"
)

//Migrate the model structs to the database
func (db connection) Migrate() {
	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Post{}, &models.Comment{}, &models.Reply{})
	if err != nil {
		loggers.Error.Fatalln(err)
	}
	
	loggers.Info.Println("Migrated tables successfully...")
}
