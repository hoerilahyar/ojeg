package db

import (
	"log"
	"ojeg/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {

	// Seed roles
	roles := []domain.Role{
		{Name: "super admin", Slug: "super-admin"},
		{Name: "admin", Slug: "admin"},
		{Name: "user", Slug: "user"},
	}

	for _, r := range roles {
		var existing domain.Role
		if err := db.Where("name = ?", r.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&r).Error; err != nil {
				log.Printf("❌ Failed seeding role: %s", r.Name)
			} else {
				log.Printf("✅ Role seeded: %s", r.Name)
			}
		}
	}

	// Seed permissions
	permissions := []domain.Permission{
		{Name: "Manage Users", Slug: "manage-users"},
		{Name: "Manage Roles", Slug: "manage-roles"},
	}

	for _, p := range permissions {
		var existing domain.Permission
		if err := db.Where("name = ?", p.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&p).Error; err != nil {
				log.Printf("❌ Failed seeding permission: %s", p.Name)
			} else {
				log.Printf("✅ Permission seeded: %s", p.Name)
			}
		}
	}

	// Seed admin user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	admin := domain.User{
		Username: "admin",
		Name:     "Administrator",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
	}

	var existingUser domain.User
	if err := db.Where("email = ?", admin.Email).First(&existingUser).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(&admin).Error; err == nil {
			log.Println("✅ Admin user created")

			// Associate role
			var adminRole domain.Role
			if err := db.Where("slug = ?", "super-admin").First(&adminRole).Error; err == nil {
				db.Model(&admin).Association("Roles").Append(&adminRole)
				log.Println("✅ Admin role assigned to user")
			}
		}
	}
}
