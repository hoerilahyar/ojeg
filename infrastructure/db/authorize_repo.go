package db

import (
	"context"
	"ojeg/internal/dto"
	"ojeg/internal/repository"

	"gorm.io/gorm"
)

type authorizeRepository struct {
	db *gorm.DB
}

func NewAuthorizeRepository(db *gorm.DB) repository.AuthorizeRepository {
	return &authorizeRepository{db: db}
}

func (r *authorizeRepository) AssignRole(ctx context.Context, dto dto.AssignRoleDTO) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO user_roles (user_id, role_id) VALUES (?, ?) ON CONFLICT DO NOTHING
	`, dto.UserID, dto.RoleID).Error
}

func (r *authorizeRepository) RevokeRole(ctx context.Context, dto dto.RevokeRoleDTO) error {
	return r.db.WithContext(ctx).Exec(`
		DELETE FROM user_roles WHERE user_id = ? AND role_id = ?
	`, dto.UserID, dto.RoleID).Error
}

func (r *authorizeRepository) AssignPermission(ctx context.Context, dto dto.AssignPermissionDTO) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?) ON CONFLICT DO NOTHING
	`, dto.RoleID, dto.PermissionID).Error
}

func (r *authorizeRepository) RevokePermission(ctx context.Context, dto dto.RevokePermissionDTO) error {
	return r.db.WithContext(ctx).Exec(`
		DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?
	`, dto.RoleID, dto.PermissionID).Error
}
