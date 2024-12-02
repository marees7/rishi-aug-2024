package validation

import (
	"blogs/pkg/models"
	"fmt"
	"net/mail"
)

func Validation(user *models.Users) error {
	//check email address
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	//check username
	if len(user.Username) <= 3 || len(user.Username) >= 20 {
		return fmt.Errorf("entered username is either too short or exceeded size limit")
	}

	//check password
	if len(user.Password) < 8 {
		return fmt.Errorf("password must contain atleast 8 characters")
	}

	return nil
}
