package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"
	"fmt"
	"net/http"
)

type AdminServices interface {
	GetUsers(users *[]models.Users, role string, limit, offset int) (int, error)
	GetUser(users *models.Users, username string, role string) (int, error)
	CreateCategory(category *models.Categories, role string) (int, error)
	UpdateCategory(category *models.Categories, categoryid int, role string) (int, error)
	DeleteCategory(category *models.Categories, categoryid int, role string) (int, error)
}

type adminService struct {
	*repositories.Repository
}

// retrieve every users records
func (repo *adminService) GetUsers(users *[]models.Users, role string, limit, offset int) (int, error) {
	if role == "admin" {
		if status, err := repo.User.GetUsers(users, limit, offset); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusOK, nil
}

// retrieve a single user records
func (repo *adminService) GetUser(users *models.Users, username string, role string) (int, error) {
	if role == "admin" {
		if status, err := repo.User.GetUser(users, username); err != nil {
			return status, err
		}
	} else {
		return http.StatusUnauthorized, fmt.Errorf("only admins can access this page")
	}
	return http.StatusOK, nil
}

// create a new category
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

// update a existing category
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

// delete a existing category
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
