package repositories

import (
	"blogs/common/dto"
	"blogs/pkg/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(limit int, offset int, name string) (*[]models.User, int64, error)
	GetUser(username string) (*models.User, *dto.ErrorResponse)
	UpdateUser(user *models.User) *dto.ErrorResponse
	DeleteUser(email string) *dto.ErrorResponse
}

type userRepository struct {
	*gorm.DB
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// retrieve every users records
func (db *userRepository) GetUsers(limit int, offset int, name string) (*[]models.User, int64, error) {
	var user []models.User
	var count int64

	//retrieve users along with comments and posts
	data := db.Model(user).Preload("Posts").Preload("Comments").Limit(limit).Offset(offset).Count(&count).Find(&user)
	if data.Error != nil {
		return nil, 0, data.Error
	}
	if name != "" {
		db.Where("name=?", name)
	}

	return &user, count, nil
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

func (db *userRepository) UpdateUser(user *models.User) *dto.ErrorResponse {
	//updates the user
	data := db.Where("email=?", user.Email).Updates(&user)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}

	return nil
}

func (db *userRepository) DeleteUser(email string) *dto.ErrorResponse {
	//deletes the user
	data := db.Where("email=?", email).Delete(&models.User{})
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}

	return nil
}
