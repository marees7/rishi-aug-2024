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
	UpdatePost(post *models.Posts, userid int, postid int) error
	CreateCategory(category *models.Categories) error
	RetrieveCategories(categories *[]models.Categories) error
}

type userRepository struct {
	*gorm.DB
}

func (db *userRepository) RetrieveUsers(user *[]models.Users) error {
	data := db.Find(&user)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	}

	return nil
}

func (db *userRepository) RetrieveSingleUser(user *models.Users, usernaem string) error {
	data := db.Where("username=?", usernaem).First(&user)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
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

func (db *userRepository) UpdatePost(post *models.Posts, userid int, postid int) error {
	var checkPost models.Posts
	data := db.Where("post_id=?", postid).First(&checkPost)
	if data.Error != nil {
		loggers.WarningLog.Println(data.Error.Error())
		return data.Error
	} else if checkPost.UserID != userid {
		loggers.WarningLog.Println("cannot update other users post")
		return fmt.Errorf("cannot update other users post")
	}

	if post.UserID == userid {
		data := db.Where("post_id=?", postid).Updates(&post)
		if data.Error != nil {
			loggers.WarningLog.Println(data.Error.Error())
			return data.Error
		}
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
