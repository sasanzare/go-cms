package models

import (
	"time"
	"gorm.io/gorm"
)

type Tag struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"size:255;not null;uniqueIndex" validate:"required,min=2,max=255"`
	Slug        string         `gorm:"size:300;uniqueIndex" validate:"omitempty,alphanumdash"`
	Status      string         `gorm:"size:20;not null;default:active" validate:"oneof=active archived"`
	CreatedAt   time.Time      `gorm:"not null;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"not null;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Posts       []Post         `gorm:"many2many:post_tags;"`
}

// Status constants for Tag
const (
	TagStatusActive   = "active"
	TagStatusArchived = "archived"
)

// BeforeCreate hook for setting default values
func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	if t.Status == "" {
		t.Status = TagStatusActive
	}
	return nil
}