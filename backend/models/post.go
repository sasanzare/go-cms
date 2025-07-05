package models

import (
	"time"
	"gorm.io/gorm"
)

type Post struct {
	ID          	uint           `gorm:"primaryKey"`
	Title       	string         `gorm:"size:255;not null" validate:"required,min=3,max=255"`
	Content     	string         `gorm:"type:text;not null" validate:"required,min=10"`
	Status      	string         `gorm:"size:20;not null;default:draft" validate:"oneof=draft published archived rejected"`
	AuthorID    	uint           `gorm:"not null"`
	ApprovedBy  	*uint
	Slug        	string         `gorm:"size:300;uniqueIndex" validate:"omitempty,alphanumdash"`
	MetaTitle       string     	   `gorm:"size:255"`
	MetaDescription string    	   `gorm:"size:500"`
	FeaturedImage   string    	   `gorm:"size:512"`
	CategoryID     	*uint
	ViewCount		uint           `gorm:"default:0"`
	CreatedAt   	time.Time      `gorm:"not null;autoCreateTime"`
	UpdatedAt  	 	time.Time      `gorm:"not null;autoUpdateTime"`
	PublishedAt 	*time.Time     `gorm:"index"`
	ApprovedAt  	*time.Time     `gorm:"index"`
	DeletedAt   	gorm.DeletedAt `gorm:"index"`
}
