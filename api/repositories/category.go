package repositories

import (
	"blogs/pkg/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategories() (*[]models.Category, error)
	UpdateCategory(category *models.Category, categoryid uuid.UUID) error
	DeleteCategory(categoryid uuid.UUID) (*models.Category, error)
}

type categoryRepository struct {
	*gorm.DB
}

func InitCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// creates a new category
func (db *categoryRepository) CreateCategory(category *models.Category) error {
	data := db.Create(&category)
	if data.Error != nil {
		return data.Error
	}
	return nil
}

// retrieve every categories available
func (db *categoryRepository) GetCategories() (*[]models.Category, error) {
	var categories []models.Category

	data := db.Find(&categories)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found")
	} else if data.Error != nil {
		return nil, data.Error
	}
	return &categories, nil
}

// update an existing category
func (db *categoryRepository) UpdateCategory(category *models.Category, categoryid uuid.UUID) error {
	var checkCategory models.Category

	//check if the category exists
	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		return data.Error
	} else if checkCategory.CategoryID != categoryid {
		return fmt.Errorf("category id not found")
	}

	//updates the category if it is the admin
	data = db.Where("category_id=?", categoryid).Updates(&category)
	if data.Error != nil {
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

// deletes the existing category
func (db *categoryRepository) DeleteCategory(categoryid uuid.UUID) (*models.Category, error) {
	var checkCategory models.Category

	//check if the record exists
	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		return nil, data.Error
	} else if checkCategory.CategoryID != categoryid {
		return nil, fmt.Errorf("category id not found")
	}

	//deletes the category if it is the admin
	data = db.Where("category_id=?", categoryid).Delete(&checkCategory)
	if data.Error != nil {
		return nil, data.Error
	} else if data.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}
	return &checkCategory, nil
}
