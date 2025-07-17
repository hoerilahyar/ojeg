package db

import (
	"context"
	"fmt"

	"ojeg/internal/domain"
	"ojeg/internal/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) ListRoles(ctx context.Context) ([]domain.Role, error) {
	var roles []domain.Role
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) GetRoleByID(ctx context.Context, id uint) (domain.Role, error) {
	var role domain.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *roleRepository) CreateRole(ctx context.Context, role *domain.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) UpdateRole(ctx context.Context, role *domain.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *roleRepository) DeleteRole(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Role{}, id).Error
}

func (r *roleRepository) GetRoleByName(ctx context.Context, name string) (*domain.Role, error) {
	fmt.Println(name)
	var role domain.Role
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
