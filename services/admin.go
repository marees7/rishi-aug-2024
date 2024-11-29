package services

import (
	"blogs/models"
	"blogs/repositories"
	"fmt"
)

type AdminServices interface {
	GetUsers(users *[]models.Users, role string) error
	GetUserByID(users *models.Users, username string, role string) error
	CreateCategory(category *models.Categories, role string) error
	UpdateCategory(category *models.Categories, categoryid int, role string) error
	DeleteCategory(category *models.Categories, categoryid int, role string) error
}

type adminService struct {
	*repositories.Repository
}

func (repo *adminService) GetUsers(users *[]models.Users, role string) error {
	if role == "admin" {
		if err := repo.User.RetrieveUsers(users); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("only admins can access this page")
	}
	return nil
}

func (repo *adminService) GetUserByID(users *models.Users, username string, role string) error {
	if role == "admin" {
		if err := repo.User.RetrieveSingleUser(users, username); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("only admins can access this page")
	}
	return nil
}

func (repo *adminService) CreateCategory(category *models.Categories, role string) error {
	if role == "admin" {
		if err := repo.User.CreateCategory(category); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("only admins can access this page")
	}
	return nil
}

func (repo *adminService) UpdateCategory(category *models.Categories, categoryid int, role string) error {
	if role == "admin" {
		if err := repo.User.UpdateCategory(category, categoryid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("only admins can access this page")
	}
	return nil
}

func (repo *adminService) DeleteCategory(category *models.Categories, categoryid int, role string) error {
	if role == "admin" {
		if err := repo.User.DeleteCategory(category, categoryid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("only admins can access this page")
	}
	return nil
}
