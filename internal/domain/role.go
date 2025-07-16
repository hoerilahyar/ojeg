package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleUser       = "user"
	RoleDriver     = "driver"
	RoleAdmin      = "admin"
	RoleSuperAdmin = "super-admin"
)

type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"type:varchar(191);uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
