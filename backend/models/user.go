package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey"`
	Username      string         `gorm:"size:50;not null;uniqueIndex" json:"username" validate:"required,alphanum,min=3,max=50"`
	FirstName 	  string 		 `gorm:"size:100;not null" json:"first_name" validate:"required,min=3,max=100"`
	LastName      string 		 `gorm:"size:100;not null" json:"last_name" validate:"required,min=3,max=100"`
	Email         string         `gorm:"size:255;uniqueIndex;not null" validate:"required,email"`
	Password      string         `gorm:"size:255;not null" validate:"required,min=8"`
	Bio           string         `gorm:"type:text"`
	Avatar        string         `gorm:"size:512"`
	Role          string         `gorm:"size:50;not null;default:user" validate:"oneof=admin editor author user"`
	Status        string         `gorm:"size:20;not null;default:active" validate:"oneof=active suspended banned"`
	LastLoginAt   *time.Time
	CreatedAt     time.Time      `gorm:"not null;autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"not null;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Posts         []Post         `gorm:"foreignKey:AuthorID"`
	ApprovedPosts []Post         `gorm:"foreignKey:ApprovedBy"`
}

// Role constants for User
const (
	UserRoleAdmin  = "admin"
	UserRoleEditor = "editor"
	UserRoleAuthor = "author"
	UserRoleUser   = "user"
)

// Status constants for User
const (
	UserStatusActive    = "active"
	UserStatusSuspended = "suspended"
	UserStatusBanned    = "banned"
)

// BeforeCreate hook for setting default values
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = UserRoleUser
	}
	if u.Status == "" {
		u.Status = UserStatusActive
	}
	return nil
}