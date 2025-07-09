package models

import (
	"time"
	"gorm.io/gorm"
)

type Post struct {
	ID          	uint           `gorm:"primaryKey"`
	Title       	string         `gorm:"size:255;not null" validate:"required,min=3,max=255"`
	Content     	string         `gorm:"type:text;not null" validate:"required,min=10"`
	Excerpt         string         `gorm:"size:500"`
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
	Category       Category   `gorm:"foreignKey:CategoryID"`
	Tags          []Tag      `gorm:"many2many:post_tags;"`
	Author      User           `gorm:"foreignKey:AuthorID"`
	Approver    User `gorm:"foreignKey:ApprovedBy"`
}

const (
    PostStatusDraft     = "draft"
    PostStatusPublished = "published"
    PostStatusArchived  = "archived"
    PostStatusRejected  = "rejected"
)

// IsPublished checks if the post is published
func (p *Post) IsPublished() bool {
    return p.Status == PostStatusPublished && p.PublishedAt != nil
}

// IsApproved checks if the post is approved
func (p *Post) IsApproved() bool {
    return p.ApprovedBy != nil && p.ApprovedAt != nil
}

// BeforeCreate hook for setting default values
func (p *Post) BeforeCreate(tx *gorm.DB) error {
    if p.Status == "" {
        p.Status = PostStatusDraft
    }
    return nil
}