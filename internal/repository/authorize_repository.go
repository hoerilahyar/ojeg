package repository

import (
	"context"
	"ojeg/internal/dto"
)

type AuthorizeRepository interface {
	AssignRole(ctx context.Context, dto dto.AssignRoleDTO) error
	RevokeRole(ctx context.Context, dto dto.RevokeRoleDTO) error
	AssignPermission(ctx context.Context, dto dto.AssignPermissionDTO) error
	RevokePermission(ctx context.Context, dto dto.RevokePermissionDTO) error
}
