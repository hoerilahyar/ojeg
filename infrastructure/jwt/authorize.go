package jwt

import (
	"net/http"
	"ojeg/internal/domain"
)

type ABACFunc func(user *domain.User, r *http.Request) bool

func Authorize(permission string, abac ABACFunc) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*domain.User)
			if !ok || user == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if !user.HasPermission(permission) {
				http.Error(w, "Forbidden (RBAC)", http.StatusForbidden)
				return
			}

			if abac != nil && !abac(user, r) {
				http.Error(w, "Forbidden (ABAC)", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
