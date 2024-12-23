package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AdminServices interface {
	GetUsers(limit, offset int, name string) (*[]models.User, int64, error)
	GetUser(username string) (*models.User, *dto.ErrorResponse)
	UpdateUser(user *models.User) *dto.ErrorResponse
	DeleteUser(email string) *dto.ErrorResponse
}

type adminService struct {
	Users repositories.UserRepository
}

func InitAdminService(user repositories.UserRepository) AdminServices {
	return &adminService{user}
}

// retrieve every users records
func (repo *adminService) GetUsers(limit, offset int, name string) (*[]models.User, int64, error) {
	return repo.Users.GetUsers(limit, offset, name)
}

// retrieve a single user records
func (repo *adminService) GetUser(username string) (*models.User, *dto.ErrorResponse) {
	return repo.Users.GetUser(username)
}

func (repo *adminService) UpdateUser(user *models.User) *dto.ErrorResponse {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: "could not update password"}
	}
	
	user.Password = string(hashedPass)

	return repo.Users.UpdateUser(user)
}

func (repo *adminService) DeleteUser(email string) *dto.ErrorResponse {
	return repo.Users.DeleteUser(email)
}
