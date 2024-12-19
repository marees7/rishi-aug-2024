package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"

	"github.com/google/uuid"
)

type CategoryServices interface {
	CreateCategory(category *models.Category) *dto.ErrorResponse
	GetCategories(limit, offset int) (*[]models.Category, *dto.ErrorResponse)
	UpdateCategory(category *models.Category, categoryid uuid.UUID) *dto.ErrorResponse
	DeleteCategory(categoryid uuid.UUID, role string) (*models.Category, *dto.ErrorResponse)
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
func (repo *userService) GetCategories(limit, offset int) (*[]models.Category, *dto.ErrorResponse) {
	offset = (offset - 1) * limit

	return repo.Category.GetCategories(limit, offset)
}

// update a existing category
func (repo *userService) UpdateCategory(category *models.Category, categoryid uuid.UUID) *dto.ErrorResponse {
	return repo.Category.UpdateCategory(category, categoryid)
}

// delete a existing category
func (repo *userService) DeleteCategory(categoryid uuid.UUID, role string) (*models.Category, *dto.ErrorResponse) {
	return repo.Category.DeleteCategory(categoryid)
}
