package db

import (
	"context"
	"ojeg/internal/domain"
	"ojeg/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindAllUser(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindUserByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.User{}, id).Error
}

func (r *userRepository) FindUserByEmailOrUsername(ctx context.Context, value string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Where("email = ? OR user_name = ?", value, value).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmailOrUsername(ctx context.Context, value string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Preload("Roles.Permissions").
		Preload("Permissions"). // if user has direct perms
		Where("email = ? OR user_name = ?", value, value).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
