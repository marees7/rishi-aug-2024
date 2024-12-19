package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"
)

type AdminServices interface {
	GetUsers(limit, offset int, name string) (*[]models.User, error)
	GetUser(username string) (*models.User, *dto.ErrorResponse)
}

type adminService struct {
	Users repositories.UserRepository
}

func InitAdminService(user repositories.UserRepository) AdminServices {
	return &adminService{user}
}

// retrieve every users records
func (repo *adminService) GetUsers(limit, offset int, name string) (*[]models.User, error) {
	offset = (offset - 1) * limit

	return repo.Users.GetUsers(limit, offset, name)
}

// retrieve a single user records
func (repo *adminService) GetUser(username string) (*models.User, *dto.ErrorResponse) {
	return repo.Users.GetUser(username)
}
