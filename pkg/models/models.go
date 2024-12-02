package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	UserID    int            `json:"user_id,omitempty" gorm:"primarykey;autoIncrement"`
	Email     string         `json:"email,omitempty" gorm:"unique;not null;"`
	Username  string         `json:"username,omitempty" gorm:"unique;not null;"`
	Password  string         `json:"password,omitempty" gorm:"not null;check:"`
	Role      string         `json:"role,omitempty" gorm:"check:role='admin' or role='user'"`
	Comments  []Comments     `json:"comments,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Posts     []Posts        `json:"posts,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Categories struct {
	CategoryID    int            `json:"category_id,omitempty" gorm:"primarykey;autoIncrement"`
	Category_name string         `json:"category_name,omitempty" gorm:"unique;not null;"`
	Description   string         `json:"description,omitempty"`
	Posts         []Posts        `json:"posts,omitempty" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

type Posts struct {
	PostID      int            `json:"post_id,omitempty" gorm:"primarykey;autoIncrement"`
	Title       string         `json:"title,omitempty" gorm:"not null;"`
	Content     string         `json:"content,omitempty" gorm:"not null;"`
	Description string         `json:"description,omitempty"`
	UserID      int            `json:"user_id,omitempty"`
	CategoryID  int            `json:"category_id,omitempty"`
	Comments    []Comments     `json:"comments,omitempty" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Comments struct {
	CommentID int            `json:"comment_id,omitempty" gorm:"primarykey;autoIncrement"`
	Content   string         `json:"content,omitempty" gorm:"not null;default:''"`
	UserID    int            `json:"user_id,omitempty"`
	PostID    int            `json:"post_id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
