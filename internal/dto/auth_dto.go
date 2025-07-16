package dto

import "ojeg/internal/domain"

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
	User      *domain.User `json:"user"`
	Token     string       `json:"token"`
	TokenType string       `json:"token_type"`
	ExpiresIn int64        `json:"expires_in"`
	IssuedAt  int64        `json:"issued_at"`
	ExpiresAt int64        `json:"expires_at"`
}
