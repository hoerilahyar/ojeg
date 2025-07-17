package domain

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Slug  string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Roles []Role `gorm:"many2many:role_permissions;" json:"roles"`
	Users []User `gorm:"many2many:user_permissions;" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
