package services

import (
	"blogs/api/repositories"
	"blogs/common/dto"
	"blogs/pkg/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthServices interface {
	Signup(user *models.User) error
	Login(login *dto.LoginRequest) (*models.User, *dto.ErrorResponse)
}

type authService struct {
	repositories.AuthRepository
}

func InitAuthService(repository repositories.AuthRepository) AuthServices {
	return &authService{repository}
}

// hashes the password and sends it to the db
func (repo *authService) Signup(user *models.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	if err := repo.AuthRepository.Signup(user); err != nil {
		return err
	}
	return nil
}

// compares the hashed password lets the user login
func (repo *authService) Login(login *dto.LoginRequest) (*models.User, *dto.ErrorResponse) {
	user, err := repo.AuthRepository.Login(login)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: "password didn't match"}
	}
	return user, nil

}
