package repository

import (
	"context"

	"ojeg/internal/domain"
)

type UserRepository interface {
	FindAllUser(ctx context.Context) ([]*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	FindUserByID(ctx context.Context, id uint) (*domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id uint) error
	FindUserByEmailOrUsername(ctx context.Context, value string) (*domain.User, error)
}
