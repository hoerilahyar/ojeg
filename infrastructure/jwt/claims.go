package jwt

import "context"

type AuthClaims struct {
	UserID uint
	Email  string
	Role   string
}

// Key used to store claims in context
type ctxKey struct{}

var claimsKey = ctxKey{}

type authContextKey string

const userContextKey authContextKey = "user_claims"

// Store user claims into context
func WithUser(ctx context.Context, claims AuthClaims) context.Context {
	return context.WithValue(ctx, userContextKey, claims)
}

// Extract user claims from context
func ExtractUser(ctx context.Context) AuthClaims {
	if val := ctx.Value(claimsKey); val != nil {
		if claims, ok := val.(AuthClaims); ok {
			return claims
		}
	}
	return AuthClaims{}
}

func GetUser(ctx context.Context) AuthClaims {
	user, _ := ctx.Value(userContextKey).(AuthClaims)
	return user
}
