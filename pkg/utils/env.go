package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Coba load dari current dir
	err := godotenv.Load(".env")
	if err != nil {
		// Coba fallback ke satu tingkat di atas (untuk unit test)
		cwd, _ := os.Getwd()
		parentEnv := filepath.Join(cwd, "../.env")

		if err := godotenv.Load(parentEnv); err != nil {
			log.Printf("Warning: .env file not found in either ./ or ../ (%v)", err)
		}
	}
}
