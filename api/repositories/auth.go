package repositories

import (
	"blogs/common/dto"
	"blogs/pkg/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Signup(*models.User) error
	Login(details *dto.LoginRequest) (*models.User, *dto.ErrorResponse)
}

type authRepository struct {
	*gorm.DB
}

func InitAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

// creates a new user
func (db *authRepository) Signup(user *models.User) error {
	data := db.Create(&user)
	if data.Error != nil {
		return data.Error
	}
	return nil
}

// login a new user
func (db *authRepository) Login(details *dto.LoginRequest) (*models.User, *dto.ErrorResponse) {
	var user models.User

	data := db.Where("email=?", details.Email).First(&user)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: "user not found"}
	} else if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}
	return &user, nil
}
