package service

import (
	"context"
	"strings"

	"ojeg/internal/user/domain"
	"ojeg/internal/user/repository"
	"ojeg/internal/user/usecase"
	"ojeg/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) usecase.UserUsecase {
	return &userService{userRepo: userRepo}
}

func (s *userService) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {

	if len(user.Password) <= 5 {
		return errors.ErrWeakPassword
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrHashFailed
	}

	user.Password = string(hashedPassword)

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(strings.ToLower(err.Error()), "unique") {
			return errors.ErrUserExists
		}
		return errors.ErrInternalError
	}
	return nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.userRepo.FindByEmail(ctx, email)
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}
