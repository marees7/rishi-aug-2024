package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"
	"fmt"
	"net/http"
)

type AdminServices interface {
	GetUsers(users *[]models.Users, role string, limit, offset int) (int, error)
	GetUserByID(users *models.Users, username string, role string) (int, error)
	CreateCategory(category *models.Categories, role string) (int, error)
	UpdateCategory(category *models.Categories, categoryid int, role string) (int, error)
	DeleteCategory(category *models.Categories, categoryid int, role string) (int, error)
}

type adminService struct {
	*repositories.Repository
}

func (repo *adminService) GetUsers(users *[]models.Users, role string, limit, offset int) (int, error) {
	if role == "admin" {
		if status, err := repo.User.RetrieveUsers(users, limit, offset); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusOK, nil
}

func (repo *adminService) GetUserByID(users *models.Users, username string, role string) (int, error) {
	if role == "admin" {
		if status, err := repo.User.RetrieveSingleUser(users, username); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusOK, nil
}

func (repo *adminService) CreateCategory(category *models.Categories, role string) (int, error) {
	if role == "admin" {
		if status, err := repo.User.CreateCategory(category); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusCreated, nil
}

func (repo *adminService) UpdateCategory(category *models.Categories, categoryid int, role string) (int, error) {
	if role == "admin" {
		if status, err := repo.User.UpdateCategory(category, categoryid); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusOK, nil
}

func (repo *adminService) DeleteCategory(category *models.Categories, categoryid int, role string) (int, error) {
	if role == "admin" {
		if status, err := repo.User.DeleteCategory(category, categoryid); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusOK, nil
}
