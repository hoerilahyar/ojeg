package main

import (
	"fmt"
	"ojeg/configs"
	"ojeg/infrastructure/db"
	"ojeg/migrations"
	"ojeg/pkg/utils"
	"os"
	"strings"
	"time"
)

func main() {
	// Check if at least one command argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ojeg [command] [optional_name]")
		return
	}

	switch os.Args[1] {
	case "make:module":
		// Ensure module name is provided
		if len(os.Args) < 3 {
			fmt.Println("⚠️  Module name is required")
			return
		}
		createModule(os.Args[2])

	case "make:migration":
		// Ensure migration name is provided
		if len(os.Args) < 3 {
			fmt.Println("⚠️  Migration name is required")
			return
		}
		createMigration(os.Args[2])

	case "migrate":
		runMigrations("migrate")

	case "migrate:refresh":
		runMigrations("refresh")

	case "migrate:rollback":
		runMigrations("rollback")

	default:
		// Unknown command
		fmt.Println("❌ Command not found:", os.Args[1])
	}
}

// createModule generates a folder structure for a new module
func createModule(name string) {
	name = strings.ToLower(name)
	fmt.Println("📦 Creating module:", name)

	dirs := []string{
		"internal/" + name + "/domain",
		"internal/" + name + "/repository",
		"internal/" + name + "/usecase",
		"internal/" + name + "/service",
		"delivery/http/handler",
	}

	// Create required directories
	for _, dir := range dirs {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	// Create empty domain file as an example
	_ = createFile("internal/" + name + "/domain/" + name + ".go")
}

// createMigration generates a migration file with timestamp
func createMigration(name string) {
	// Normalize name
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("⚠️  Migration name is required")
		return
	}

	// Auto-expand shorthand: "Role" → "create_role_table"
	lower := strings.ToLower(name)
	if !strings.HasPrefix(lower, "create_") && !strings.Contains(lower, "_") {
		name = fmt.Sprintf("create_%s_table", lower)
	}

	// Prepare naming variants
	timestamp := time.Now().Format("20060102150405")
	snake := utils.ToSnakeCase(lower)
	camel := utils.ToCamelCase(name)

	// File path
	filename := fmt.Sprintf("migrations/%s_%s.go", timestamp, snake)

	// Ensure migrations dir
	if err := os.MkdirAll("migrations", os.ModePerm); err != nil {
		fmt.Println("❌ Failed to create migrations dir:", err)
		return
	}

	// Migration file template
	template := fmt.Sprintf(`package migrations

import (
	"ojeg/internal/%s/domain"
	"gorm.io/gorm"
)

func Up_%s(db *gorm.DB) error {
	return db.AutoMigrate(&domain.%s{})
}

func Down_%s(db *gorm.DB) error {
	return db.Migrator().DropTable(&domain.%s{})
}
`, lower, camel, camel, camel, camel, camel)

	if err := os.WriteFile(filename, []byte(template), 0644); err != nil {
		fmt.Println("❌ Failed to write migration file:", err)
		return
	}

	// Registry update
	registryPath := "migrations/registry.go"
	regEntry := fmt.Sprintf(`	{
		Name: "%s",
		Up:   Up_%s,
		Down: Down_%s,
	},
`, snake, camel, camel)

	regContent, err := os.ReadFile(registryPath)
	if err != nil {
		fmt.Println("❌ Failed to read registry.go:", err)
		return
	}

	newRegistry := insertMigrationToRegistry(string(regContent), regEntry)
	if err := os.WriteFile(registryPath, []byte(newRegistry), 0644); err != nil {
		fmt.Println("❌ Failed to update registry.go:", err)
		return
	}

	fmt.Println("✅ Migration created:", filename)
	fmt.Println("✅ registry.go updated with:", camel)
}

// createFile creates a basic Go file with the appropriate package name
func createFile(path string) error {
	// Check if file already exists
	if _, err := os.Stat(path); err == nil {
		fmt.Println("⚠️  File already exists:", path)
		return nil
	}

	// Generate package name based on the directory
	parts := strings.Split(path, "/")
	packageName := parts[len(parts)-2]

	// Write a simple Go file template
	content := fmt.Sprintf("package %s\n\n", packageName)

	return os.WriteFile(path, []byte(content), 0644)
}

func runMigrations(mode string) {
	// Load config
	cfg := configs.LoadConfig()

	// Connect to DB
	conn, err := db.NewDB(db.DBConfig{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})

	if err != nil {
		fmt.Println("❌ Failed to connect to DB:", err)
		return
	}
	gormDB := conn

	switch mode {
	case "migrate":
		fmt.Println("🚀 Running migrations...")
		for _, m := range migrations.MigrationList {
			fmt.Println("⬆️  Applying:", m.Name)
			if err := m.Up(gormDB); err != nil {
				fmt.Println("❌ Failed:", m.Name, err)
				return
			}
			fmt.Println("✅ Success:", m.Name)
		}

	case "rollback":
		fmt.Println("↩️  Rolling back migrations...")
		for i := len(migrations.MigrationList) - 1; i >= 0; i-- {
			m := migrations.MigrationList[i]
			fmt.Println("🔽 Reverting:", m.Name)
			if err := m.Down(gormDB); err != nil {
				fmt.Println("❌ Rollback failed:", m.Name, err)
				return
			}
			fmt.Println("✅ Rolled back:", m.Name)
		}

	case "refresh":
		fmt.Println("🔄 Refreshing migrations...")
		runMigrations("rollback")
		runMigrations("migrate")

	default:
		fmt.Println("❌ Unknown migration mode:", mode)
	}
}

func insertMigrationToRegistry(content, newEntry string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.Contains(line, "var MigrationList = []Migration{") {
			// Find end of list
			j := i + 1
			for ; j < len(lines); j++ {
				if strings.TrimSpace(lines[j]) == "}" {
					break
				}
			}
			// Insert before closing brace
			lines = append(lines[:j], append(strings.Split(newEntry, "\n"), lines[j:]...)...)
			break
		}
	}
	return strings.Join(lines, "\n")
}
