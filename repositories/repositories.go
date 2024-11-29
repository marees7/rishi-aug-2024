package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB   *gorm.DB
	Auth AuthRepository
	User UserRepository
}

func GetRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB:   db,
		Auth: &authRepository{db},
		User: &userRepository{db},
	}
}
