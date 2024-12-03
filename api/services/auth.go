package services

import (
	"blogs/common/helpers"
	"blogs/pkg/models"
	"blogs/api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthServices interface {
	Signup(user *models.Users) error
	Login(login *helpers.LoginRequest) (*models.Users, error)
}

type authService struct {
	*repositories.Repository
}

//hashes the password and sends it to the db
func (repo *authService) Signup(user *models.Users) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	if err := repo.Auth.RegisterUser(user); err != nil {
		return err
	}
	return nil
}

//compares the hashed password lets the user login
func (repo *authService) Login(login *helpers.LoginRequest) (*models.Users, error) {
	user, err := repo.Auth.LoginUser(login)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return nil, err
	}
	return user, nil

}
