package service

import (
	"context"
	"ojeg/internal/domain"
	"ojeg/internal/repository"
	"ojeg/internal/usecase"
)

type PermissionDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type permissionService struct {
	repo repository.PermissionRepository
}

func PermissionService(repo repository.PermissionRepository) usecase.PermissionUsecase {
	return &permissionService{repo: repo}
}

func (s *permissionService) ListPermissions(ctx context.Context) ([]domain.Permission, error) {
	return s.repo.ListPermissions(ctx)
}

func (s *permissionService) CreatePermission(ctx context.Context, permission *domain.Permission) error {
	return s.repo.CreatePermission(ctx, permission)
}

func (s *permissionService) GetPermissionByID(ctx context.Context, id uint) (domain.Permission, error) {
	return s.repo.GetPermissionByID(ctx, id)
}

func (s *permissionService) UpdatePermission(ctx context.Context, permission *domain.Permission) error {
	return s.repo.UpdatePermission(ctx, permission)
}

func (s *permissionService) DeletePermission(ctx context.Context, id uint) error {
	return s.repo.DeletePermission(ctx, id)
}
