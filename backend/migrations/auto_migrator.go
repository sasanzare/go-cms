package migrations

import (
	"fmt"
	"time"
	"log"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type AutoMigrator struct {
	db      *gorm.DB
	models  []interface{}
	verbose bool
}

func NewAutoMigrator(db *gorm.DB, verbose bool) *AutoMigrator {
	return &AutoMigrator{
		db:      db,
		verbose: verbose,
	}
}


func (am *AutoMigrator) AddModel(model interface{}) {
	am.models = append(am.models, model)
}

func (am *AutoMigrator) AddModels(models ...interface{}) {
	am.models = append(am.models, models...)
}


func (am *AutoMigrator) Run() error {
	if am.verbose {
		log.Println("Starting auto migration process...")
		log.Printf("Models to migrate: %d\n", len(am.models))
	}

	if err := am.db.AutoMigrate(&MigrationRecord{}); err != nil {
		return fmt.Errorf("failed to migrate migrations table: %v", err)
	}

	for _, model := range am.models {
		modelName := getModelName(model)
		migrationID := generateMigrationID(modelName)

		var record MigrationRecord
		if err := am.db.Where("id = ?", migrationID).First(&record).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return fmt.Errorf("failed to check migration record: %v", err)
			}

			if am.verbose {
				log.Printf("Migrating model: %s\n", modelName)
			}

			if err := am.db.AutoMigrate(model); err != nil {
				return fmt.Errorf("failed to auto migrate model %s: %v", modelName, err)
			}

			if err := am.db.Create(&MigrationRecord{
				ID:        migrationID,
				CreatedAt: time.Now(),
			}).Error; err != nil {
				return fmt.Errorf("failed to create migration record: %v", err)
			}

			if am.verbose {
				log.Printf("Successfully migrated model: %s\n", modelName)
			}
		}
	}

	if am.verbose {
		log.Println("Auto migration completed successfully")
	}
	return nil
}


func generateMigrationID(modelName string) string {
	return fmt.Sprintf("auto_%s_%d", strings.ToLower(modelName), time.Now().Unix())
}

func getModelName(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}