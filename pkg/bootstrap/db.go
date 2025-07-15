package bootstrap

import (
	"log"
	"os"

	"ojeg/configs"
	"ojeg/infrastructure/db"
)

// InitDB initializes and returns the database connection using loaded config
func InitDB() *db.DB {
	// Load app config (based on APP_ENV)
	cfg := configs.LoadConfig()

	// Use .env for sensitive values
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("❌ DB_PASSWORD not set in environment")
	}

	// Connect using infrastructure/db
	database, err := db.NewDB(db.DBConfig{
		Driver:   cfg.DB.Driver,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: password,
		Name:     cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Fatalf("❌ DB connect failed: %v", err)
	}

	log.Println("✅ Database initialized")
	return database
}
