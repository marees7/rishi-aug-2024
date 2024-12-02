package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	Auth AuthRepository
	User UserRepository
}

func GetAuthRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth: &authRepository{db},
	}
}

func GetUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: &userRepository{db},
	}
}
