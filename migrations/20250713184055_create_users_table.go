package migrations

import (
	"ojeg/internal/domain"

	"gorm.io/gorm"
)

// Up_create_users_table runs the migration to create the "users" table
func Up_create_users_table(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{})
}

// Down_create_users_table drops the "users" table
func Down_create_users_table(db *gorm.DB) error {
	return db.Migrator().DropTable(&domain.User{})
}
