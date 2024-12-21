package repositories

import (
	"blogs/common/dto"
	"blogs/pkg/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Signup(*models.User) *dto.ErrorResponse
	Login(details *dto.LoginRequest) (*models.User, *dto.ErrorResponse)
}

type authRepository struct {
	*gorm.DB
}

func InitAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

// creates a new user
func (db *authRepository) Signup(user *models.User) *dto.ErrorResponse {
	//check if the user already exists
	data := db.Where("email=? OR username = ?", user.Email, user.Username).First(&user)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		//create the user
		data = db.Create(&user)
		if data.Error != nil {
			return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
		}
		
		return nil
	} else if data.RowsAffected > 0 {
		return &dto.ErrorResponse{Status: http.StatusConflict, Error: "user already exists"}
	} else {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}
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
