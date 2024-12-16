package repositories

import (
	"blogs/pkg/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post, postid uuid.UUID, role string) error
	DeletePost(userid uuid.UUID, postid uuid.UUID, role string) (*models.Post, error)
	GetPosts(startdate string, enddate string, postid uuid.UUID, title string) (*[]models.Post, error)
}

type postRepository struct {
	*gorm.DB
}

func InitpostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

// create a new post
func (db *postRepository) CreatePost(post *models.Post) error {
	//creates a new post
	data := db.Create(&post)
	if data.Error != nil {
		return data.Error
	}
	return nil
}

// update a existing post
func (db *postRepository) UpdatePost(post *models.Post, postid uuid.UUID, role string) error {
	var checkPost models.Post

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		return data.Error
	} else if checkPost.UserID != post.UserID && role != "admin" {
		return fmt.Errorf("cannot update other users post")
	}

	//updates the record if the user created it or if it is the admin
	if post.UserID == checkPost.UserID || role == "admin" {
		data := db.Where("post_id=?", postid).Updates(&post)
		if data.Error != nil {
			return data.Error
		} else if data.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}
	}
	return nil
}

// delete a existing post
func (db *postRepository) DeletePost(userid uuid.UUID, postid uuid.UUID, role string) (*models.Post, error) {
	var checkPost models.Post

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		return nil, data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		return nil, fmt.Errorf("cannot delete other users post")
	}

	//deletes the record if the user created it or if it is the admin
	if checkPost.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Delete(&checkPost)
		if data.Error != nil {
			return nil, data.Error
		} else if data.RowsAffected == 0 {
			return nil, fmt.Errorf("no rows affected")
		}
	}
	return &checkPost, nil
}

// retrieve every users posts using either date or post id
func (db *postRepository) GetPosts(startdate string, enddate string, postid uuid.UUID, title string) (*[]models.Post, error) {
	var post []models.Post
	//check whether to use date filter or post id
	if startdate != "" && enddate != "" && postid != uuid.Nil {
		data := db.Where("created_at BETWEEN ? AND ? AND post_id= ?", startdate, enddate, postid).Preload("Comments").Find(&post)
		if data.RowsAffected == 0 {
			return nil, fmt.Errorf("no records found")
		}
		if data.Error != nil {
			return nil, data.Error
		}
	} else if startdate != "" && enddate != "" {
		data := db.Where("created_at BETWEEN ? AND ?", startdate, enddate).Preload("Comments").Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return nil, fmt.Errorf("no records found")
			} else {
				return nil, data.Error
			}
		}
	} else if postid != uuid.Nil {
		data := db.Where("post_id=?", postid).Preload("Comments").Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return nil, fmt.Errorf("no records found")
			} else {
				return nil, data.Error
			}
		}
	} else if title != "" {
		data := db.Where("title LIKE '%' || ? || '%' ", title).Preload("Comments").Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return nil, fmt.Errorf("no records found")
			} else {
				return nil, data.Error
			}
		}
	} else {
		data := db.Preload("Comments").Find(&post)
		if data.RowsAffected == 0 {
			return nil, fmt.Errorf("no records found")
		}
		if data.Error != nil {
			return nil, data.Error
		}
	}
	return &post, nil
}
