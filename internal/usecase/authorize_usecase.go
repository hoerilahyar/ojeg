package usecase

import (
	"context"
	"ojeg/internal/dto"
)

type AuthorizeUsecase interface {
	AssignRole(ctx context.Context, dto dto.AssignRoleDTO) error
	RevokeRole(ctx context.Context, dto dto.RevokeRoleDTO) error
	AssignPermission(ctx context.Context, dto dto.AssignPermissionDTO) error
	RevokePermission(ctx context.Context, dto dto.RevokePermissionDTO) error
}
