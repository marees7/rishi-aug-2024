package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"
)

type AdminServices interface {
	GetUsers(role string, limit, offset int) (*[]models.User, error)
	GetUser(username string, role string) (*models.User, error)
}

type adminService struct {
	Users repositories.UserRepository
}

func InitAdminService(user repositories.UserRepository) AdminServices {
	return &adminService{user}
}

// retrieve every users records
func (repo *adminService) GetUsers(role string, limit, offset int) (*[]models.User, error) {
	users, err := repo.Users.GetUsers(limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// retrieve a single user records
func (repo *adminService) GetUser(username string, role string) (*models.User, error) {
	users, err := repo.Users.GetUser(username)
	if err != nil {
		return nil, err
	}
	return users, nil
}
