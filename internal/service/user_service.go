package service

import (
	"context"
	"strings"

	"ojeg/internal/domain"
	"ojeg/internal/repository"
	"ojeg/internal/usecase"
	"ojeg/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func UserService(userRepo repository.UserRepository) usecase.UserUsecase {
	return &userService{userRepo: userRepo}
}

func (s *userService) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return s.userRepo.FindAllUser(ctx)
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

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(strings.ToLower(err.Error()), "unique") {
			return errors.ErrUserExists
		}
		return errors.ErrInternalError
	}
	return nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.userRepo.FindUserByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.userRepo.FindUserByEmail(ctx, email)
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.DeleteUser(ctx, id)
}
