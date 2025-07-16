package usecase

import (
	"context"
	"ojeg/internal/domain"
)

type RoleUsecase interface {
	ListRoles(ctx context.Context) ([]domain.Role, error)
	GetRoleByID(ctx context.Context, id uint) (domain.Role, error)
	CreateRole(ctx context.Context, role *domain.Role) error
	UpdateRole(ctx context.Context, role *domain.Role) error
	DeleteRole(ctx context.Context, id uint) error
}
