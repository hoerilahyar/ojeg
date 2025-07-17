package jwt

import (
	"fmt"
	"net/http"
	"ojeg/internal/domain"
	"ojeg/pkg/errors"
	"ojeg/pkg/response"
)

// ABACFunc is a function that returns true if access is allowed
type ABACFunc func(user *domain.User, r *http.Request) bool

func Authorize(permission string, abac ABACFunc) func(w http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, ok := r.Context().Value("user").(*domain.User)
			if !ok || user == nil {
				fmt.Println(r.Context().Value("user"))
				response.Error(w, errors.ErrUnauthorized)
				return
			}

			// Super admin shortcut
			if user.HasRole("super-admin") {
				next.ServeHTTP(w, r)
				return
			}

			// RBAC check
			if !user.HasPermission(permission) {
				response.Error(w, errors.ErrForbiddenPermission)
				return
			}

			// ABAC check
			if abac != nil && !abac(user, r) {
				response.Error(w, errors.ErrForbiddenAccessDenied)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
