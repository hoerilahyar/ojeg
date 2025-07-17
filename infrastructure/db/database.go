package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

type DB = gorm.DB

// NewDB creates and returns a new GORM database connection based on config.
func NewDB(config DBConfig) (*DB, error) {
	var dialector gorm.Dialector

	switch config.Driver {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			config.Host, config.User, config.Password, config.Name, config.Port, config.SSLMode,
		)
		dialector = postgres.Open(dsn)

	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci",
			config.User, config.Password, config.Host, config.Port, config.Name,
		)
		dialector = mysql.Open(dsn)

	case "sqlite":
		dialector = sqlite.Open(config.Name) // Use "ojeg.db" or ":memory:"

	default:
		return nil, fmt.Errorf("❌ unsupported DB driver: %s", config.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to open DB: %w", err)
	}

	// AutoMigrate domain models (you can extract this later into its own function)
	// err = db.AutoMigrate(
	// 	&domain.User{},
	// 	&domain.Role{},
	// 	&domain.Permission{},
	// 	&domain.UserRole{},
	// 	&domain.UserPermission{},
	// 	&domain.RolePermission{},
	// )
	// if err != nil {
	// 	log.Fatalf("❌ AutoMigrate failed: %v", err)
	// }

	// Seed(db)

	log.Println("✅ Connected using", config.Driver)
	return db, nil
}
