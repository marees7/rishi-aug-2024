package validation

import (
	"blogs/common/constants"
	"blogs/pkg/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// validates the user fields
func ValidateUser(user *models.User) error {
	//check email address
	validate := validator.New()
	err := validate.Struct(user)
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

	//check role
	if user.Role != constants.AdminRole && user.Role != "user" {
		return fmt.Errorf("role must be either admin or user")
	}

	return nil
}

// Validates the category fields
func ValidateCategory(category *models.Category) error {
	//check category name
	if category.CategoryName == "" {
		return fmt.Errorf("category name cannot be empty")
	}

	//check description
	if category.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	return nil
}

// Validate the posts fields
func ValidatePost(post *models.Post) error {
	//check title
	if post.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	//check content
	if post.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}

	//check description
	if post.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	return nil
}

// Validate the comments fields
func ValidateComment(comment *models.Comment) error {
	//check content
	if comment.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}

	return nil
}

// Validate the reply fields
func ValidateReply(reply *models.Reply) error {
	//check content
	if reply.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}

	return nil
}

// Check role
func ValidateRole(role string) bool {
	return role == constants.AdminRole
}
