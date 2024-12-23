package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// contains the user details
type User struct {
	UserID    uuid.UUID      `json:"user_id,omitempty" gorm:"type:uuid;primary_key"`
	Email     string         `json:"email,omitempty" validate:"required,email" gorm:"unique;not null;"`
	Name      string         `json:"name,omitempty" gorm:"not null;default:'anonymous'"`
	Username  string         `json:"username,omitempty" gorm:"unique;not null;"`
	Password  string         `json:"password,omitempty" gorm:"not null;"`
	Role      string         `json:"role,omitempty" gorm:"check:role='admin' or role='user'"`
	Comments  []Comment      `json:"comments,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Replies   []Reply        `json:"replies,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Posts     []Post         `json:"posts,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// contains category details
type Category struct {
	CategoryID   uuid.UUID      `json:"category_id,omitempty" gorm:"type:uuid;primary_key"`
	CategoryName string         `json:"category_name,omitempty" gorm:"unique;not null;"`
	Description  string         `json:"description,omitempty"`
	Posts        []Post         `json:"posts,omitempty" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

// contains post details
type Post struct {
	PostID      uuid.UUID      `json:"post_id,omitempty" gorm:"type:uuid;primary_key"`
	Title       string         `json:"title,omitempty" gorm:"not null;"`
	Content     string         `json:"content,omitempty" gorm:"not null;"`
	Description string         `json:"description,omitempty"`
	UserID      uuid.UUID      `json:"user_id,omitempty" gorm:"type:uuid"`
	CategoryID  uuid.UUID      `json:"category_id,omitempty" gorm:"type:uuid"`
	Comments    []Comment      `json:"comments,omitempty" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

// contains the comment details
type Comment struct {
	CommentID uuid.UUID      `json:"comment_id,omitempty" gorm:"type:uuid;primary_key"`
	Content   string         `json:"content,omitempty" gorm:"not null;default:''"`
	UserID    uuid.UUID      `json:"user_id,omitempty" gorm:"type:uuid"`
	PostID    uuid.UUID      `json:"post_id,omitempty" gorm:"type:uuid"`
	Replies   []Reply        `json:"replies,omitempty" gorm:"foreignKey:CommentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// contains the reply details
type Reply struct {
	ReplyID   uuid.UUID      `json:"reply_id,omitempty" gorm:"type:uuid;primary_key"`
	Content   string         `json:"content,omitempty" gorm:"not null;"`
	UserID    uuid.UUID      `json:"user_id,omitempty" gorm:"type:uuid"`
	CommentID uuid.UUID      `json:"comment_id,omitempty" gorm:"type:uuid"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// assign uuid before insert a new row
func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.UserID = uuid.New()
	return nil
}

// assign uuid before insert a new row
func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.PostID = uuid.New()
	return nil
}

// assign uuid before insert a new row
func (comment *Comment) BeforeCreate(tx *gorm.DB) error {
	comment.CommentID = uuid.New()
	return nil
}

// assign uuid before insert a new row
func (category *Category) BeforeCreate(tx *gorm.DB) error {
	category.CategoryID = uuid.New()
	return nil
}

// assign uuid before insert a new row
func (reply *Reply) BeforeCreate(tx *gorm.DB) error {
	reply.ReplyID = uuid.New()
	return nil
}
