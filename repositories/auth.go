package repositories

import (
	"blogs/helpers"
	"blogs/loggers"
	"blogs/models"
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

func (db *authRepository) RegisterUser(user *models.Users) error {
	data := db.Create(&user)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}

	return nil
}

func (db *authRepository) LoginUser(details *helpers.LoginRequest) (*models.Users, error) {
	var user models.Users

	data := db.Where("email=?", details.Email).First(&user)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return nil, data.Error
	} else if user.UserID == 0 {
		loggers.WarningLog.Println("user not found")
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}
