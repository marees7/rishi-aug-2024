package repositories

import (
	"blogs/initializers"

	"gorm.io/gorm"
)

type Repository struct {
	gormDB *gorm.DB
	auth   AuthRepository
	user   UserRepository
}

func GetRepository() *Repository {
	db := initializers.GetDB()
	return &Repository{
		gormDB: db,
		auth:   &authRepository{db},
		user:   &userRepository{db},
	}
}
