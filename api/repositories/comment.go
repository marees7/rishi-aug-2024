package repositories

import (
	"blogs/pkg/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	UpdateComment(comment *models.Comment, commentid uuid.UUID, role string) error
	GetComment(postid uuid.UUID, content string) (*[]models.Comment, error)
	DeleteComment(userid uuid.UUID, commentid uuid.UUID, role string) (*models.Comment, error)
}

type commentRepository struct {
	*gorm.DB
}

func InitCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

// create a new comment
func (db *commentRepository) CreateComment(comment *models.Comment) error {
	data := db.Create(&comment)
	if data.Error != nil {
		return data.Error
	}
	return nil
}

// retrieve comments using post id
func (db *commentRepository) GetComment(postid uuid.UUID, content string) (*[]models.Comment, error) {
	var checkComment models.Comment
	var comment []models.Comment

	//retrieves the comment
	if content != "" && postid == uuid.Nil {
		//get the comments using content
		data := db.Where("content LIKE '%' || ? || '%'", content).Find(&comment)
		if errors.Is(data.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		} else if data.Error != nil {
			return nil, data.Error
		}
		return &comment, nil
	} else if postid != uuid.Nil && content == "" {
		//check if the post exists
		data := db.Where("post_id=?", postid).First(&checkComment)
		if data.Error != nil {
			return nil, data.Error
		} else if checkComment.PostID != postid {
			return nil, fmt.Errorf("post id not found")
		}

		//get the comments using postid
		data = db.Where("post_id=?", postid).Find(&comment)
		if errors.Is(data.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		} else if data.Error != nil {
			return nil, data.Error
		}
		return &comment, nil
	} else {
		data := db.Find(&comment)
		if data.Error != nil {
			return nil, data.Error
		}
		return &comment, nil
	}
}

// updates the existing comment
func (db *commentRepository) UpdateComment(comment *models.Comment, commentid uuid.UUID, role string) error {
	var checkComment models.Comment

	//check if the record exists and if the user can access it
	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		return data.Error
	} else if checkComment.UserID != comment.UserID && role != "admin" {
		return fmt.Errorf("cannot update other users comment")
	}

	//updates the comment if the user created it or if it is the admin
	if checkComment.UserID == comment.UserID || role == "admin" {
		data = db.Where("comment_id=?", commentid).Updates(&comment)
		if data.Error != nil {
			return data.Error
		} else if data.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}
	}
	return nil
}

// deletes the existing comment
func (db *commentRepository) DeleteComment(userid uuid.UUID, commentid uuid.UUID, role string) (*models.Comment, error) {
	var checkComment models.Comment

	//check if the record exists and if the user can access it
	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		return nil, data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		return nil, fmt.Errorf("cannot delete other users comment")
	}

	//deletes the record if the user created it or if it is the admin
	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Delete(&checkComment)
		if data.Error != nil {
			return nil, data.Error
		} else if data.RowsAffected == 0 {
			return nil, fmt.Errorf("no rows affected")
		}
	}
	return &checkComment, nil
}
