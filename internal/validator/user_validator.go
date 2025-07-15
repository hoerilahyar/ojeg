package validator

import (
	"ojeg/internal/domain"
	"ojeg/pkg/errors"
	"strings"
)

func ValidateRegisterInput(req *domain.RegisterRequest) *errors.AppError {
	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.UserName) == "" ||
		strings.TrimSpace(req.Email) == "" ||
		strings.TrimSpace(req.Password) == "" {
		return errors.ErrMissingFields
	}

	if !strings.Contains(req.Email, "@") {
		return errors.ErrInvalidEmail
	}

	if len(req.Password) < 6 {
		return errors.ErrWeakPassword
	}

	return nil
}
