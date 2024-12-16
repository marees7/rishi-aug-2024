package validation

import (
	"blogs/pkg/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// validates the user fields
func UserValidation(user *models.User) error {
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
	if user.Role != "admin" && user.Role != "user" {
		return fmt.Errorf("role must be either admin or user")
	}

	return nil
}

// Validates the category fields
func CategoryValidation(category *models.Category) error {
	//check category name
	if category.Category_name == "" {
		return fmt.Errorf("category name cannot be empty")
	}

	//check description
	if category.Description == "" {
		return fmt.Errorf("descroption cannot be empty")
	}

	return nil
}

// Validate the posts fields
func PostsValidation(post *models.Post) error {
	//check category name
	if post.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	//check content
	if post.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}

	//check description
	if post.Description == "" {
		return fmt.Errorf("descroption cannot be empty")
	}

	return nil
}

// Validate the comments fields
func CommentValidation(comment *models.Comment) error {
	//check content
	if comment.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}

	return nil
}
