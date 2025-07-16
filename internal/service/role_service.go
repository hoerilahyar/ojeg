package service

import (
	"context"
	"ojeg/internal/domain"
	"ojeg/internal/repository"
	"ojeg/internal/usecase"
)

type RoleDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type roleService struct {
	repo repository.RoleRepository
}

func RoleService(repo repository.RoleRepository) usecase.RoleUsecase {
	return &roleService{repo: repo}
}

func (s *roleService) ListRoles(ctx context.Context) ([]domain.Role, error) {
	return s.repo.ListRoles(ctx)
}

func (s *roleService) CreateRole(ctx context.Context, role *domain.Role) error {
	return s.repo.CreateRole(ctx, role)
}

func (s *roleService) GetRoleByID(ctx context.Context, id uint) (domain.Role, error) {
	return s.repo.GetRoleByID(ctx, id)
}

func (s *roleService) UpdateRole(ctx context.Context, role *domain.Role) error {
	return s.repo.UpdateRole(ctx, role)
}

func (s *roleService) DeleteRole(ctx context.Context, id uint) error {
	return s.repo.DeleteRole(ctx, id)
}
