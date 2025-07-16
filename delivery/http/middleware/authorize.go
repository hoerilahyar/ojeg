package middleware

import (
	"context"
	"net/http"
	"strings"

	jwtInfra "ojeg/infrastructure/jwt"
	"ojeg/internal/repository"
	"ojeg/pkg/errors"
	"ojeg/pkg/response"
)

func JWTMiddleware(jwtService jwtInfra.JWTService, userRepo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				response.Error(w, errors.ErrUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			token, claims, err := jwtService.ValidateToken(tokenStr)
			if err != nil || !token.Valid {
				response.Error(w, errors.ErrUnauthorized)
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				response.Error(w, errors.ErrUnauthorized)
				return
			}
			userID := uint(userIDFloat)

			user, err := userRepo.FindUserByID(r.Context(), userID)
			if err != nil {
				response.Error(w, errors.ErrUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
