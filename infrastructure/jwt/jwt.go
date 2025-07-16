package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTService defines the interface for generating JWT tokens
type JWTService interface {
	GenerateToken(userID uint, email string, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new instance of JWTService
func NewJWTService(secretKey string, issuer string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

// GenerateToken generates a JWT token with custom claims
func (j *jwtService) GenerateToken(userID uint, email string, role string) (string, error) {
	expireStr := os.Getenv("JWT_EXPIRE_HOURS")
	expireHours := 72 // default
	if expireStr != "" {
		if h, err := strconv.Atoi(expireStr); err == nil {
			expireHours = h
		}
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"iss":     j.issuer,
		"exp":     time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, jwt.ErrInvalidKey
}
