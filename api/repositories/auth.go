package repositories

import (
	"blogs/common/helpers"
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(*models.Users) error
	LoginUser(details *helpers.LoginRequest) (*models.Users, error)
}

type authRepository struct {
	*gorm.DB
}

// creates a new user
func (db *authRepository) RegisterUser(user *models.Users) error {
	data := db.Create(&user)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return data.Error
	}

	return nil
}

// login a new user
func (db *authRepository) LoginUser(details *helpers.LoginRequest) (*models.Users, error) {
	var user models.Users

	data := db.Where("email=?", details.Email).First(&user)
	if user.UserID == 0 {
		loggers.Warn.Println("user not found")
		return nil, fmt.Errorf("user not found")
	} else if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return nil, data.Error
	}

	return &user, nil
}
