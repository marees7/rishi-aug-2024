package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"
)

type AdminServices interface {
	GetUsers(limit, page int) (*[]models.User, error)
	GetUser(username string) (*models.User, error)
}

type adminService struct {
	Users repositories.UserRepository
}

func InitAdminService(user repositories.UserRepository) AdminServices {
	return &adminService{user}
}

// retrieve every users records
func (repo *adminService) GetUsers(limit, page int) (*[]models.User, error) {
	offset := (page - 1) * limit
	
	return repo.Users.GetUsers(limit, offset)
}

// retrieve a single user records
func (repo *adminService) GetUser(username string) (*models.User, error) {
	return repo.Users.GetUser(username)
}
