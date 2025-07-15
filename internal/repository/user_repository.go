package repository

import (
	"context"

	"ojeg/internal/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
	FindByEmailOrUsername(ctx context.Context, value string) (*domain.User, error)
}
