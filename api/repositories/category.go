package repositories

import (
	"blogs/common/dto"
	"blogs/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) *dto.ErrorResponse
	GetCategories(limit, offset int) (*[]models.Category, *dto.ErrorResponse)
	UpdateCategory(category *models.Category, categoryID uuid.UUID) *dto.ErrorResponse
	DeleteCategory(categoryID uuid.UUID) (*models.Category, *dto.ErrorResponse)
}

type categoryRepository struct {
	*gorm.DB
}

func InitCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// creates a new category
func (db *categoryRepository) CreateCategory(category *models.Category) *dto.ErrorResponse {
	//check if the category already exists
	data := db.Where("category_name=?", category.CategoryName).First(&category)
	if data.RowsAffected > 0 {
		return &dto.ErrorResponse{Status: http.StatusConflict, Error: "category already exists"}
	}

	//create the category
	data = db.Create(&category)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}
	return nil
}

// retrieve every categories available
func (db *categoryRepository) GetCategories(limit, offset int) (*[]models.Category, *dto.ErrorResponse) {
	var categories []models.Category

	data := db.Limit(limit).Offset(offset).Find(&categories)
	if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: data.Error.Error()}
	}
	return &categories, nil
}

// update an existing category
func (db *categoryRepository) UpdateCategory(category *models.Category, categoryID uuid.UUID) *dto.ErrorResponse {
	var categoryData models.Category

	//check if the category exists
	data := db.Where("category_id=?", categoryID).First(&categoryData)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	}

	//updates the category if it is the admin
	data = db.Where("category_id=?", categoryID).Updates(&category)
	if data.Error != nil {
		return &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}
	category.CategoryID = categoryData.CategoryID

	return nil
}

// deletes the existing category
func (db *categoryRepository) DeleteCategory(categoryID uuid.UUID) (*models.Category, *dto.ErrorResponse) {
	var categoryData models.Category

	//check if the record exists
	data := db.Where("category_id=?", categoryID).First(&categoryData)
	if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: "category not found"}
	}

	//deletes the category if it is the admin
	data = db.Where("category_id=?", categoryID).Delete(&categoryData)
	if data.Error != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusNotFound, Error: data.Error.Error()}
	} else if data.RowsAffected == 0 {
		return nil, &dto.ErrorResponse{Status: http.StatusNotModified, Error: "no changes were made"}
	}
	return &categoryData, nil
}
