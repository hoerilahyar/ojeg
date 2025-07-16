package db

import (
	"context"

	"ojeg/internal/domain"
	"ojeg/internal/repository"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) ListPermissions(ctx context.Context) ([]domain.Permission, error) {
	var permissions []domain.Permission
	if err := r.db.WithContext(ctx).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *permissionRepository) GetPermissionByID(ctx context.Context, id uint) (domain.Permission, error) {
	var permission domain.Permission
	if err := r.db.WithContext(ctx).First(&permission, id).Error; err != nil {
		return permission, err
	}
	return permission, nil
}

func (r *permissionRepository) CreatePermission(ctx context.Context, permission *domain.Permission) error {
	return r.db.WithContext(ctx).Create(permission).Error
}

func (r *permissionRepository) UpdatePermission(ctx context.Context, permission *domain.Permission) error {
	return r.db.WithContext(ctx).Save(permission).Error
}

func (r *permissionRepository) DeletePermission(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Permission{}, id).Error
}
