package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserName  string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	Role      string         `gorm:"type:varchar(50);default:'user'" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
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
