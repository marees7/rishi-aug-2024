package repositories

import (
	"blogs/loggers"
	"blogs/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	RetrieveUsers(user *[]models.Users) error
	RetrieveSingleUser(user *models.Users, usernaem string) error
	CreatePost(post *models.Posts, userid int) error
	UpdatePost(post *models.Posts, userid int, postid int, role string) error
	DeletePost(post *models.Posts, userid int, postid int, role string) error
	RetrievePostByDate(post *[]models.Posts, startdate string, enddate string) error
	RetrievePostWithID(post *models.Posts, postid int) error
	CreateCategory(category *models.Categories) error
	RetrieveCategories(categories *[]models.Categories) error
	UpdateCategory(category *models.Categories, categoryid int) error
	DeleteCategory(category *models.Categories, categoryid int) error
	CreateComment(comment *models.Comments, userid int) error
	RetrieveComment(comment *[]models.Comments, postid int) error
	UpdateComment(comment *models.Comments, userid int, commentid int, role string) error
	DeleteComment(comment *models.Comments, userid int, commentid int, role string) error
}

type userRepository struct {
	*gorm.DB
}

func (db *userRepository) RetrieveUsers(user *[]models.Users) error {
	data := db.Find(&user)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (db *userRepository) RetrieveSingleUser(user *models.Users, username string) error {
	data := db.Where("username=?", username).First(&user)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (db *userRepository) CreatePost(post *models.Posts, userid int) error {
	post.UserID = userid

	data := db.Create(&post)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}
	return nil
}

func (db *userRepository) UpdatePost(post *models.Posts, userid int, postid int, role string) error {
	var checkPost models.Posts

	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("cannot update other users post")
		return fmt.Errorf("cannot update other users post")
	}

	if post.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Updates(&post)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return data.Error
		} else if data.RowsAffected == 0 {
			return fmt.Errorf("could not update post id:%d", postid)
		}
	}
	return nil
}

func (db *userRepository) DeletePost(post *models.Posts, userid int, postid int, role string) error {
	var checkPost models.Posts

	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkPost.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("cannot delete other users post")
		return fmt.Errorf("cannot delete other users post")
	}

	if post.UserID == userid || role == "admin" {
		data := db.Where("post_id=?", postid).Delete(&post)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return data.Error
		} else if data.RowsAffected == 0 {
			return fmt.Errorf("could not delete post id:%d", postid)
		}
	}
	return nil
}

func (db *userRepository) RetrievePostByDate(post *[]models.Posts, startdate string, enddate string) error {
	data := db.Where("created_at BETWEEN ? AND ?", startdate, enddate).Find(&post)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (db *userRepository) RetrievePostWithID(post *models.Posts, postid int) error {
	data := db.Where("post_id=?", postid).Find(&post)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (db *userRepository) CreateCategory(category *models.Categories) error {
	data := db.Create(&category)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}
	return nil
}

func (db *userRepository) RetrieveCategories(categories *[]models.Categories) error {
	data := db.Find(&categories)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}
	return nil
}

func (db *userRepository) UpdateCategory(category *models.Categories, categoryid int) error {
	var checkCategory models.Categories

	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkCategory.CategoryID != categoryid {
		loggers.WarningLog.Println("category id not found")
		return fmt.Errorf("category id not found")
	}

	data = db.Where("category_id=?", categoryid).Updates(&category)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (db *userRepository) DeleteCategory(category *models.Categories, categoryid int) error {
	var checkCategory models.Categories

	data := db.Where("category_id=?", categoryid).First(&checkCategory)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkCategory.CategoryID != categoryid {
		loggers.WarningLog.Println("category id not found")
		return fmt.Errorf("category id not found")
	}

	data = db.Where("category_id=?", categoryid).Delete(&category)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if data.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (db *userRepository) CreateComment(comment *models.Comments, userid int) error {
	comment.UserID = userid

	data := db.Create(&comment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}
	return nil
}

func (db *userRepository) RetrieveComment(comment *[]models.Comments, postid int) error {
	data := db.Where("post_id=?", postid).Find(&comment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}
	return nil
}

func (db *userRepository) UpdateComment(comment *models.Comments, userid int, commentid int, role string) error {
	var checkComment models.Comments

	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("Cannot update other users comment")
		return fmt.Errorf("cannot update other users comment")
	}

	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Updates(&comment)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return data.Error
		} else if data.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}
	}
	return nil
}

func (db *userRepository) DeleteComment(comment *models.Comments, userid int, commentid int, role string) error {
	var checkComment models.Comments

	data := db.Where("comment_id=?", commentid).First(&checkComment)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkComment.UserID != userid && role != "admin" {
		loggers.WarningLog.Println("Cannot delete other users comment")
		return fmt.Errorf("cannot delete other users comment")
	}

	if checkComment.UserID == userid || role == "admin" {
		data = db.Where("comment_id=?", commentid).Delete(&comment)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return data.Error
		} else if data.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}
	}
	return nil
}
