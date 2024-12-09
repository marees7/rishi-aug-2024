package repositories

import (
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(user *[]models.Users, limit int, offset int) (int, error)
	GetUser(user *models.Users, usernaem string) (int, error)
	CreatePost(post *models.Posts, userid int) (int, error)
	UpdatePost(post *models.Posts, userid int, postid int, role string) (int, error)
	DeletePost(post *models.Posts, userid int, postid int, role string) (int, error)
	GetPost(post *[]models.Posts, startdate string, enddate string, postid int) (int, error)
	CreateCategory(category *models.Categories) (int, error)
	GetCategories(categories *[]models.Categories) (int, error)
	UpdateCategory(category *models.Categories, categoryid int) (int, error)
	DeleteCategory(category *models.Categories, categoryid int) (int, error)
	CreateComment(comment *models.Comments, userid, postid int) (int, error)
	GetComment(comment *[]models.Comments, postid int) (int, error)
	UpdateComment(comment *models.Comments, userid int, commentid int, role string) (int, error)
	DeleteComment(comment *models.Comments, userid int, commentid int, role string) (int, error)
}

type userRepository struct {
	*gorm.DB
}

// retrieve every users records
func (db *userRepository) GetUsers(user *[]models.Users, limit int, offset int) (int, error) {
	//retrieve users along with comments and posts
	data := db.Preload("Posts").Preload("Comments").Limit(limit).Offset(offset).Find(&user)
	if data.RowsAffected == 0 {
		return http.StatusOK, nil
	}
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}

	return http.StatusOK, nil
}

// retrieve a single user record
func (db *userRepository) GetUser(user *models.Users, username string) (int, error) {
	//retrieve a single user record along with comments and posts
	data := db.Preload("Posts").Preload("Comments").Where("username=?", username).First(&user)
	if data.RowsAffected == 0 {
		return http.StatusOK, fmt.Errorf("no records found")
	}
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusOK, nil
}

// create a new post
func (db *userRepository) CreatePost(post *models.Posts, userid int) (int, error) {
	post.UserID = userid

	//creates a new post
	data := db.Create(&post)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusCreated, nil
}

// update a existing post
func (db *userRepository) UpdatePost(post *models.Posts, userid int, postid int, role string) (int, error) {
	var checkPost models.Posts

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		loggers.Warn.Println("cannot update other users post")
		return http.StatusUnauthorized, fmt.Errorf("cannot update other users post")
	}

	//updates the record if the user created it or if it is the admin
	if post.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Updates(&post)
		if data.Error != nil {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("could not update post id:%d", postid)
		}
	}
	return http.StatusOK, nil
}

// delete a existing post
func (db *userRepository) DeletePost(post *models.Posts, userid int, postid int, role string) (int, error) {
	var checkPost models.Posts

	//check if the record exists and if the user can access it
	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		loggers.Warn.Println("cannot delete other users post")
		return http.StatusUnauthorized, fmt.Errorf("cannot delete other users post")
	}

	//deletes the record if the user created it or if it is the admin
	if post.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Delete(&post)
		if data.Error != nil {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("no rows affected")
		}
	}
	return http.StatusOK, nil
}

// retrieve every users posts using either date or post id
func (db *userRepository) GetPost(post *[]models.Posts, startdate string, enddate string, postid int) (int, error) {
	//check whether to use date filter or post id
	if startdate != "" && enddate != "" && postid != 0 {
		data := db.Where("created_at BETWEEN ? AND ? AND post_id= ?", startdate, enddate, postid).Preload("Comments").Find(&post)
		if data.RowsAffected == 0 {
			return http.StatusOK, nil
		}
		if data.Error != nil {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	} else if startdate != "" && enddate != "" {
		data := db.Where("created_at BETWEEN ? AND ?", startdate, enddate).Preload("Comments").Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return http.StatusNotFound, fmt.Errorf("no records found")
			} else {
				loggers.Warn.Println(data.Error.Error())
				return http.StatusInternalServerError, data.Error
			}
		}
	} else if postid != 0 {
		data := db.Where("post_id=?", postid).Preload("Comments").Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return http.StatusNotFound, fmt.Errorf("no records found")
			} else {
				loggers.Warn.Println(data.Error.Error())
				return http.StatusInternalServerError, data.Error
			}
		}
	} else {
		data := db.Preload("Comments").Find(&post)
		if data.RowsAffected == 0 {
			return http.StatusOK, nil
		}
		if data.Error != nil {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	}
	return http.StatusOK, nil
}

// creates a new category
func (db *userRepository) CreateCategory(category *models.Categories) (int, error) {
	data := db.Create(&category)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusCreated, nil
}

// retrieve every categories available
func (db *userRepository) GetCategories(categories *[]models.Categories) (int, error) {
	data := db.Find(&categories)
	if data.RowsAffected == 0 {
		return http.StatusOK, nil
	}
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusOK, nil
}

// update an existing category
func (db *userRepository) UpdateCategory(category *models.Categories, categoryid int) (int, error) {
	var checkCategory models.Categories

	//check if the category exists
	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkCategory.CategoryID != categoryid {
		loggers.Warn.Println("category id not found")
		return http.StatusNotFound, fmt.Errorf("category id not found")
	}

	//updates the category if it is the admin
	data = db.Where("category_id=?", categoryid).Updates(&category)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	} else if data.RowsAffected == 0 {
		return http.StatusNotModified, fmt.Errorf("no rows affected")
	}
	return http.StatusOK, nil
}

// deletes the existing category
func (db *userRepository) DeleteCategory(category *models.Categories, categoryid int) (int, error) {
	var checkCategory models.Categories

	//check if the record exists
	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkCategory.CategoryID != categoryid {
		loggers.Warn.Println("category id not found")
		return http.StatusNotFound, fmt.Errorf("category id not found")
	}

	//deletes the category if it is the admin
	data = db.Where("category_id=?", categoryid).Delete(&category)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	} else if data.RowsAffected == 0 {
		return http.StatusNotModified, fmt.Errorf("no rows affected")
	}
	return http.StatusOK, nil
}

// create a new comment
func (db *userRepository) CreateComment(comment *models.Comments, userid, postid int) (int, error) {
	comment.UserID = userid
	comment.PostID = postid

	data := db.Create(&comment)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusCreated, nil
}

// retrieve comments using post id
func (db *userRepository) GetComment(comment *[]models.Comments, postid int) (int, error) {
	var checkComment models.Comments

	//check if the record exists
	data := db.Where("post_id=?", postid).First(&checkComment)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkComment.PostID != postid {
		loggers.Warn.Println("post id not found")
		return http.StatusNotFound, fmt.Errorf("post id not found")
	}

	//retrieves the comment
	data = db.Where("post_id=?", postid).Find(&comment)
	if data.Error != nil {
		if data.RowsAffected == 0 {
			return http.StatusNotFound, fmt.Errorf("no records found")
		} else {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	}
	return http.StatusOK, nil
}

// updates the existing comment
func (db *userRepository) UpdateComment(comment *models.Comments, userid int, commentid int, role string) (int, error) {
	var checkComment models.Comments

	//check if the record exists and if the user can access it
	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		loggers.Warn.Println("Cannot update other users comment")
		return http.StatusUnauthorized, fmt.Errorf("cannot update other users comment")
	}

	//updates the comment if the user created it or if it is the admin
	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Updates(&comment)
		if data.Error != nil {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("no rows affected")
		}
	}
	return http.StatusOK, nil
}

// deletes the existing comment
func (db *userRepository) DeleteComment(comment *models.Comments, userid int, commentid int, role string) (int, error) {
	var checkComment models.Comments

	//check if the record exists and if the user can access it
	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		loggers.Warn.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		loggers.Warn.Println("Cannot delete other users comment")
		return http.StatusUnauthorized, fmt.Errorf("cannot delete other users comment")
	}

	//deletes the record if the user created it or if it is the admin
	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Delete(&comment)
		if data.Error != nil {
			loggers.Warn.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("no rows affected")
		}
	}
	return http.StatusOK, nil
}
