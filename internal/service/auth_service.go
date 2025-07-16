package service

import (
	"context"
	"ojeg/pkg/errors"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"ojeg/infrastructure/jwt"
	"ojeg/internal/domain"
	"ojeg/internal/repository"
	"ojeg/internal/usecase"
)

type authService struct {
	userRepo   repository.UserRepository
	jwtService jwt.JWTService
}

// AuthService creates a new AuthUsecase implementation
func AuthService(userRepo repository.UserRepository, jwtService jwt.JWTService) usecase.AuthUsecase {
	return &authService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

// Register creates a new user with hashed password
func (a *authService) Register(ctx context.Context, req *domain.RegisterRequest) error {
	// Check if user already exists
	existingUser, _ := a.userRepo.FindUserByEmail(ctx, req.Email)
	if existingUser != nil {
		return errors.ErrUserExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		UserName: req.UserName,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Roles: []domain.Role{
			{
				Name: "user",
			},
		},
	}

	return a.userRepo.CreateUser(ctx, user)
}

// Login authenticates the user and returns a JWT token
func (a *authService) Login(ctx context.Context, req *domain.AuthRequest) (domain.LoginResponse, error) {
	user, err := a.userRepo.FindUserByEmailOrUsername(ctx, req.UserName)
	if err != nil {
		return domain.LoginResponse{}, errors.ErrInvalidCredentials
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return domain.LoginResponse{}, errors.ErrInvalidCredentials
	}

	var roleNames []string
	for _, r := range user.Roles {
		roleNames = append(roleNames, r.Name)
	}
	roleString := strings.Join(roleNames, ",")

	// Generate JWT token
	token, err := a.jwtService.GenerateToken(user.ID, user.Email, roleString)
	if err != nil {
		return domain.LoginResponse{}, errors.ErrInvalidCredentials
	}

	expireStr := os.Getenv("JWT_EXPIRE_HOURS")
	expireHours := 24 // default to 24 hours
	if expireStr != "" {
		if h, err := strconv.Atoi(expireStr); err == nil {
			expireHours = h
		}
	}

	now := time.Now()
	exp := now.Add(time.Duration(expireHours) * time.Hour)

	return domain.LoginResponse{
		User:      user,
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: int64(24 * 3600),
		IssuedAt:  now.Unix(),
		ExpiresAt: exp.Unix(),
	}, nil
}
