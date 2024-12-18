package repositories

import (
	"blogs/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(limit int, offset int) (*[]models.User, error)
	GetUser(username string) (*models.User, error)
}

type userRepository struct {
	*gorm.DB
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// retrieve every users records
func (db *userRepository) GetUsers(limit int, offset int) (*[]models.User, error) {
	var user []models.User
	//retrieve users along with comments and posts
	data := db.Preload("Posts").Preload("Comments").Limit(limit).Offset(offset).Find(&user)
	if data.Error != nil {
		return nil, data.Error
	}
	return &user, nil
}

// retrieve a single user record
func (db *userRepository) GetUser(username string) (*models.User, error) {
	var user models.User
	//retrieve a single user record along with comments and posts
	data := db.Preload("Posts").Preload("Comments").Where("username=?", username).First(&user)
	if data.Error != nil {
		return nil, data.Error
	}
	return &user, nil
}
