package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"ojeg/configs/schema"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// LoadConfig loads configuration from YAML and overrides with environment variables (if available).
func LoadConfig() *schema.Config {
	// Load .env file (optional but recommended)
	_ = godotenv.Load(".env")

	// Determine environment: development, production, etc.
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "default"
	}

	// Load config.yaml content
	content, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatalf("❌ Failed to read config.yaml: %v", err)
	}

	// Unmarshal YAML into map of environments
	var full map[string]*schema.Config
	if err := yaml.Unmarshal(content, &full); err != nil {
		log.Fatalf("❌ Failed to parse config.yaml: %v", err)
	}

	// Select config for current environment
	cfg, ok := full[env]
	if !ok {
		log.Fatalf("❌ Config not found for APP_ENV=%s", env)
	}

	// Override DB config from environment variables (with fallback to YAML)
	cfg.DB.Host = getEnvOrDefault("DB_HOST", cfg.DB.Host)

	if portStr := os.Getenv("DB_PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			cfg.DB.Port = p
		}
	}

	cfg.DB.Driver = getEnvOrDefault("DB_DRIVER", cfg.DB.Driver)
	cfg.DB.User = getEnvOrDefault("DB_USER", cfg.DB.User)
	cfg.DB.Password = getEnvOrDefault("DB_PASSWORD", cfg.DB.Password)
	cfg.DB.Name = getEnvOrDefault("DB_NAME", cfg.DB.Name)
	cfg.DB.SSLMode = getEnvOrDefault("DB_SSLMODE", cfg.DB.SSLMode)

	// JWT Secret
	cfg.JWTSecret = getEnvOrDefault("JWT_SECRET", cfg.JWTSecret)

	fmt.Printf("✅ Loaded config for: %s\n", env)
	return cfg
}

// getEnvOrDefault returns the environment variable if set, otherwise the fallback value.
func getEnvOrDefault(envVar string, fallback string) string {
	val := os.Getenv(envVar)
	if val == "" {
		return fallback
	}
	return val
}
