package migrations

import (
	"gorm.io/gorm"
)

type Create_role_table struct {
	gorm.Model
}

func Up_Create_role_table(db *gorm.DB) error {
	return db.AutoMigrate(&Create_role_table{})
}

func Down_Create_role_table(db *gorm.DB) error {
	return db.Migrator().DropTable(&Create_role_table{})
}
