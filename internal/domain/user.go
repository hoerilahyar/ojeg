package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserName    string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Email       string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password"`
	Roles       []Role         `gorm:"many2many:user_roles;" json:"roles"`
	Permissions []Permission   `gorm:"many2many:user_permissions;" json:"permissions"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type AuthRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User      *User  `json:"user"`
	Token     string `json:"token"`
	TokenType string `json:"token_type"` // e.g. "Bearer"
	ExpiresIn int64  `json:"expires_in"` // e.g. 86400 (24 hours in seconds)
	IssuedAt  int64  `json:"issued_at"`  // Unix timestamp
	ExpiresAt int64  `json:"expires_at"` // Unix timestamp
}

type UserRole struct {
	UserID    uint
	RoleID    uint
	CreatedAt time.Time
}

type UserPermission struct {
	UserID       uint
	PermissionID uint
	CreatedAt    time.Time
}

func (u *User) GetAllPermissions() []string {
	permMap := make(map[string]bool)

	// Permissions from roles
	for _, role := range u.Roles {
		for _, p := range role.Permissions {
			permMap[p.Name] = true
		}
	}

	// Direct permissions (optional)
	for _, p := range u.Permissions {
		permMap[p.Name] = true
	}

	var permissions []string
	for name := range permMap {
		permissions = append(permissions, name)
	}
	return permissions
}

func (u *User) HasPermission(name string) bool {
	for _, perm := range u.GetAllPermissions() {
		if perm == name {
			return true
		}
	}
	return false
}
