package services

import (
	"blogs/api/repositories"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type CategoryServices interface {
	GetCategories() (*[]models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category, categoryid uuid.UUID) error
	DeleteCategory(categoryid uuid.UUID, role string) (*models.Category, error)
}

type userService struct {
	Category repositories.CategoryRepository
}

func InitCategoryService(category repositories.CategoryRepository) CategoryServices {
	return &userService{category}
}

// retrieve every categories
func (repo *userService) GetCategories() (*[]models.Category, error) {
	category, err := repo.Category.GetCategories()
	if err != nil {
		return nil, err
	}
	return category, nil
}

// create a new category
func (repo *userService) CreateCategory(category *models.Category) error {
	if err := repo.Category.CreateCategory(category); err != nil {
		return err
	}
	return nil
}

// update a existing category
func (repo *userService) UpdateCategory(category *models.Category, categoryid uuid.UUID) error {
	if err := repo.Category.UpdateCategory(category, categoryid); err != nil {
		return err
	}
	return nil
}

// delete a existing category
func (repo *userService) DeleteCategory(categoryid uuid.UUID, role string) (*models.Category, error) {
	category, err := repo.Category.DeleteCategory(categoryid)
	if err != nil {
		return nil, err
	}
	return category, nil
}
