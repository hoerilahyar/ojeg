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
	Slug        string       `gorm:"type:varchar(191);uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
	Users       []User       `gorm:"many2many:user_roles;" json:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
