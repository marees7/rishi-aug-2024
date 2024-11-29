package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	UserID    int        `gorm:"primarykey;autoIncrement"`
	Email     string     `json:"email" gorm:"unique;not null;"`
	Username  string     `json:"username" gorm:"unique;not null;"`
	Password  string     `json:"password" gorm:"not null;check:"`
	Role      string     `json:"role" gorm:"check:role='admin' or role='user'"`
	Comments  []Comments `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Posts     []Posts    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time  `gorm:"autoCreateTime;"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt
}

type Categories struct {
	CategoryID    int       `gorm:"primarykey;autoIncrement"`
	Category_name string    `json:"category_name" gorm:"unique;not null;"`
	Description   string    `json:"description"`
	Posts         []Posts   `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time `gorm:"autoCreateTime;"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime;"`
	DeletedAt     gorm.DeletedAt
}

type Posts struct {
	PostID      int        `gorm:"primarykey;autoIncrement"`
	Title       string     `json:"title" gorm:"not null;"`
	Content     string     `json:"content" gorm:"not null;"`
	Description string     `json:"description"`
	UserID      int        `json:"user_id"`
	CategoryID  int        `json:"category_id"`
	Comments    []Comments `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time  `gorm:"autoCreateTime;"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime;"`
	DeletedAt   gorm.DeletedAt
}

type Comments struct {
	CommentID int       `gorm:"primarykey;autoIncrement"`
	Content   string    `json:"content" gorm:"not null;default:''"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt
}
