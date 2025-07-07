package migrations

import (
	"log"
	"time"

	"gorm.io/gorm"
	"github.com/sasanzare/go-cms/models"
)

type MigrationRecord struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
}


func InitAutoMigrations(db *gorm.DB) error {
	migrator := NewAutoMigrator(db, true)

	migrator.AddModels(
		// &models.User{},
		&models.Post{},
		// &models.Category{},
	)

	if err := migrator.Run(); err != nil {
		log.Printf("Migration error: %v", err)
		return err
	}

	return nil
}