package repositories

import (
	"blogs/pkg/loggers"
	"blogs/pkg/models"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	RetrieveUsers(user *[]models.Users, limit int, offset int) (int, error)
	RetrieveSingleUser(user *models.Users, usernaem string) (int, error)
	CreatePost(post *models.Posts, userid int) (int, error)
	UpdatePost(post *models.Posts, userid int, postid int, role string) (int, error)
	DeletePost(post *models.Posts, userid int, postid int, role string) (int, error)
	RetrievePost(post *[]models.Posts, startdate string, enddate string, postid int) (int, error)
	CreateCategory(category *models.Categories) (int, error)
	RetrieveCategories(categories *[]models.Categories) (int, error)
	UpdateCategory(category *models.Categories, categoryid int) (int, error)
	DeleteCategory(category *models.Categories, categoryid int) (int, error)
	CreateComment(comment *models.Comments, userid, postid int) (int, error)
	RetrieveComment(comment *[]models.Comments, postid int) (int, error)
	UpdateComment(comment *models.Comments, userid int, commentid int, role string) (int, error)
	DeleteComment(comment *models.Comments, userid int, commentid int, role string) (int, error)
}

type userRepository struct {
	*gorm.DB
}

func (db *userRepository) RetrieveUsers(user *[]models.Users, limit int, offset int) (int, error) {
	data := db.Preload("Posts").Preload("Comments").Limit(limit).Offset(offset).Find(&user)
	if data.Error != nil {
		if data.RowsAffected == 0 {
			return http.StatusNotFound, fmt.Errorf("no records found")
		} else {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) RetrieveSingleUser(user *models.Users, username string) (int, error) {
	data := db.Preload("Posts").Preload("Comments").Where("username=?", username).First(&user)
	if data.Error != nil {
		if data.RowsAffected == 0 {
			return http.StatusNotFound, fmt.Errorf("no records found")
		} else {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) CreatePost(post *models.Posts, userid int) (int, error) {
	post.UserID = userid

	data := db.Create(&post)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusCreated, nil
}

func (db *userRepository) UpdatePost(post *models.Posts, userid int, postid int, role string) (int, error) {
	var checkPost models.Posts

	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("cannot update other users post")
		return http.StatusUnauthorized, fmt.Errorf("cannot update other users post")
	}

	if post.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Updates(&post)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("could not update post id:%d", postid)
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) DeletePost(post *models.Posts, userid int, postid int, role string) (int, error) {
	var checkPost models.Posts

	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("cannot delete other users post")
		return http.StatusUnauthorized, fmt.Errorf("cannot delete other users post")
	}

	if post.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Delete(&post)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("no rows affected")
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) RetrievePost(post *[]models.Posts, startdate string, enddate string, postid int) (int, error) {
	if startdate != "" && enddate != "" && postid != 0 {
		data := db.Where("created_at BETWEEN ? AND ? AND post_id= ?", startdate, enddate, postid).Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return http.StatusNotFound, fmt.Errorf("no records found")
			} else {
				loggers.WarningLog.Println(data.Error.Error())
				return http.StatusInternalServerError, data.Error
			}
		}
	} else if startdate != "" && enddate != "" {
		data := db.Where("created_at BETWEEN ? AND ?", startdate, enddate).Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return http.StatusNotFound, fmt.Errorf("no records found")
			} else {
				loggers.WarningLog.Println(data.Error.Error())
				return http.StatusInternalServerError, data.Error
			}
		}
	} else if postid != 0 {
		data := db.Where("post_id=?", postid).Find(&post)
		if data.Error != nil {
			if data.RowsAffected == 0 {
				return http.StatusNotFound, fmt.Errorf("no records found")
			} else {
				loggers.WarningLog.Println(data.Error.Error())
				return http.StatusInternalServerError, data.Error
			}
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) CreateCategory(category *models.Categories) (int, error) {
	data := db.Create(&category)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusCreated, nil
}

func (db *userRepository) RetrieveCategories(categories *[]models.Categories) (int, error) {
	data := db.Find(&categories)
	if data.Error != nil {
		if data.RowsAffected == 0 {
			return http.StatusNotFound, fmt.Errorf("no records found")
		} else {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) UpdateCategory(category *models.Categories, categoryid int) (int, error) {
	var checkCategory models.Categories

	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkCategory.CategoryID != categoryid {
		loggers.WarningLog.Println("category id not found")
		return http.StatusNotFound, fmt.Errorf("category id not found")
	}

	data = db.Where("category_id=?", categoryid).Updates(&category)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	} else if data.RowsAffected == 0 {
		return http.StatusNotModified, fmt.Errorf("no rows affected")
	}
	return http.StatusOK, nil
}

func (db *userRepository) DeleteCategory(category *models.Categories, categoryid int) (int, error) {
	var checkCategory models.Categories

	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkCategory.CategoryID != categoryid {
		loggers.WarningLog.Println("category id not found")
		return http.StatusNotFound, fmt.Errorf("category id not found")
	}

	data = db.Where("category_id=?", categoryid).Delete(&category)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	} else if data.RowsAffected == 0 {
		return http.StatusNotModified, fmt.Errorf("no rows affected")
	}
	return http.StatusOK, nil
}

func (db *userRepository) CreateComment(comment *models.Comments, userid, postid int) (int, error) {
	comment.UserID = userid
	comment.PostID = postid

	data := db.Create(&comment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusInternalServerError, data.Error
	}
	return http.StatusCreated, nil
}

func (db *userRepository) RetrieveComment(comment *[]models.Comments, postid int) (int, error) {
	var checkComment models.Comments

	data := db.Where("post_id=?", postid).First(&checkComment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkComment.PostID != postid {
		loggers.WarningLog.Println("post id not found")
		return http.StatusNotFound, fmt.Errorf("post id not found")
	}

	data = db.Where("post_id=?", postid).Find(&comment)
	if data.Error != nil {
		if data.RowsAffected == 0 {
			return http.StatusNotFound, fmt.Errorf("no records found")
		} else {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) UpdateComment(comment *models.Comments, userid int, commentid int, role string) (int, error) {
	var checkComment models.Comments

	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("Cannot update other users comment")
		return http.StatusUnauthorized, fmt.Errorf("cannot update other users comment")
	}

	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Updates(&comment)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("no rows affected")
		}
	}
	return http.StatusOK, nil
}

func (db *userRepository) DeleteComment(comment *models.Comments, userid int, commentid int, role string) (int, error) {
	var checkComment models.Comments

	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return http.StatusNotFound, data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("Cannot delete other users comment")
		return http.StatusUnauthorized, fmt.Errorf("cannot delete other users comment")
	}

	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Delete(&comment)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return http.StatusInternalServerError, data.Error
		} else if data.RowsAffected == 0 {
			return http.StatusNotModified, fmt.Errorf("no rows affected")
		}
	}
	return http.StatusOK, nil
}
