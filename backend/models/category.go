package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
	ID          	uint           `gorm:"primaryKey"`
	Name        	string         `gorm:"size:255;not null" validate:"required,min=3,max=255"`
	Description 	string         `gorm:"type:text"`
	Status      	string         `gorm:"size:20;not null;default:draft" validate:"oneof=draft published archived"`
	Slug        	string         `gorm:"size:300;uniqueIndex" validate:"omitempty,alphanumdash"`
	MetaTitle       string        `gorm:"size:255"`
	MetaDescription string        `gorm:"size:500"`
	FeaturedImage   string        `gorm:"size:512"`
	ParentID     	*uint          `gorm:"index"`
	DeletedAt   	gorm.DeletedAt `gorm:"index"`
	Parent      	*Category      `gorm:"foreignKey:ParentID"`
	Posts       	[]Post         `gorm:"foreignKey:CategoryID"`
	CreatedAt   	time.Time      `gorm:"not null;autoCreateTime"`
	UpdatedAt   	time.Time      `gorm:"not null;autoUpdateTime"`
	PublishedAt 	*time.Time     `gorm:"index"`
}

// Status constants for Category
const (
	CategoryStatusDraft     = "draft"
	CategoryStatusPublished = "published"
	CategoryStatusArchived  = "archived"
)

// IsPublished checks if the category is published
func (c *Category) IsPublished() bool {
	return c.Status == CategoryStatusPublished && c.PublishedAt != nil
}

// BeforeCreate hook for setting default values
func (c *Category) BeforeCreate(tx *gorm.DB) error {
	if c.Status == "" {
		c.Status = CategoryStatusDraft
	}
	return nil
}