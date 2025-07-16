package usecase

import (
	"context"
	"ojeg/internal/domain"
)

type PermissionUsecase interface {
	ListPermissions(ctx context.Context) ([]domain.Permission, error)
	GetPermissionByID(ctx context.Context, id uint) (domain.Permission, error)
	CreatePermission(ctx context.Context, permission *domain.Permission) error
	UpdatePermission(ctx context.Context, permission *domain.Permission) error
	DeletePermission(ctx context.Context, id uint) error
}
