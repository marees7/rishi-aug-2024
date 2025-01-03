package services

import (
	"github.com/marees7/rishi-aug-2024/api/repositories"
	"github.com/marees7/rishi-aug-2024/common/dto"
	"github.com/marees7/rishi-aug-2024/pkg/models"

	"github.com/google/uuid"
)

type CategoryServices interface {
	CreateCategory(category *models.Category) *dto.ErrorResponse
	GetCategories(limit, offset int) (*[]models.Category, *dto.ErrorResponse, int64)
	UpdateCategory(category *models.Category, categoryID uuid.UUID) *dto.ErrorResponse
	DeleteCategory(categoryID uuid.UUID) *dto.ErrorResponse
}

type userService struct {
	Category repositories.CategoryRepository
}

func InitCategoryService(category repositories.CategoryRepository) CategoryServices {
	return &userService{category}
}

// create a new category
func (repo *userService) CreateCategory(category *models.Category) *dto.ErrorResponse {
	return repo.Category.CreateCategory(category)
}

// retrieve every categories
func (repo *userService) GetCategories(limit, offset int) (*[]models.Category, *dto.ErrorResponse, int64) {
	return repo.Category.GetCategories(limit, offset)
}

// update a existing category
func (repo *userService) UpdateCategory(category *models.Category, categoryID uuid.UUID) *dto.ErrorResponse {
	return repo.Category.UpdateCategory(category, categoryID)
}

// delete a existing category
func (repo *userService) DeleteCategory(categoryID uuid.UUID) *dto.ErrorResponse {
	return repo.Category.DeleteCategory(categoryID)
}
