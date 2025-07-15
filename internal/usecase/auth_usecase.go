package usecase

import (
	"context"
	"ojeg/internal/domain"
)

type AuthUsecase interface {
	Register(ctx context.Context, req *domain.RegisterRequest) error
	Login(ctx context.Context, req *domain.AuthRequest) (domain.LoginResponse, error) // returns JWT token
}
