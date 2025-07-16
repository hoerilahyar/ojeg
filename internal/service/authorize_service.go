package service

import (
	"context"
	"ojeg/internal/dto"
	"ojeg/internal/repository"
	"ojeg/internal/usecase"
)

type authorizeService struct {
	repo repository.AuthorizeRepository
}

func AuthorizeService(repo repository.AuthorizeRepository) usecase.AuthorizeUsecase {
	return &authorizeService{repo: repo}
}

func (s *authorizeService) AssignRole(ctx context.Context, dto dto.AssignRoleDTO) error {
	return s.repo.AssignRole(ctx, dto)
}

func (s *authorizeService) RevokeRole(ctx context.Context, dto dto.RevokeRoleDTO) error {
	return s.repo.RevokeRole(ctx, dto)
}

func (s *authorizeService) AssignPermission(ctx context.Context, dto dto.AssignPermissionDTO) error {
	return s.repo.AssignPermission(ctx, dto)
}

func (s *authorizeService) RevokePermission(ctx context.Context, dto dto.RevokePermissionDTO) error {
	return s.repo.RevokePermission(ctx, dto)
}
