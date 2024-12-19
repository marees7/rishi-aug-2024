package repositories

import (
	"blogs/common/dto"
	"blogs/pkg/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(limit int, offset int, name string) (*[]models.User, error)
	GetUser(username string) (*models.User, *dto.ErrorResponse)
}

type userRepository struct {
	*gorm.DB
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// retrieve every users records
func (db *userRepository) GetUsers(limit int, offset int, name string) (*[]models.User, error) {
	var user []models.User
	//retrieve users along with comments and posts
	data := db.Preload("Posts").Preload("Comments").Limit(limit).Offset(offset).Find(&user)
	if data.Error != nil {
		return nil, data.Error
	}
	if name != "" {
		db.Where("name=?", name)
	}
	return &user, nil
}

// retrieve a single user record
func (db *userRepository) GetUser(username string) (*models.User, *dto.ErrorResponse) {
	var user models.User
	//retrieve a single user record along with comments and posts
	data := db.Preload("Posts").Preload("Comments").Where("username=?", username).First(&user)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: "user not found"}
	} else if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}
	return &user, nil
}
