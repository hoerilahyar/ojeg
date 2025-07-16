package http

import (
	"net/http"
	"os"

	"ojeg/delivery/http/middleware"
	"ojeg/delivery/http/registry"
	"ojeg/infrastructure/db"
	"ojeg/infrastructure/jwt"

	"github.com/gorilla/mux"
)

func NewRouter(r *registry.HandlerRegistry, database *db.DB) http.Handler {
	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	secret := os.Getenv("JWT_SECRET")
	issuer := "ojeg"

	jwtService := jwt.NewJWTService(secret, issuer)
	userRepo := db.NewUserRepository(database)

	// Inisialisasi jwtMiddleware
	jwtMiddleware := middleware.JWTMiddleware(jwtService, userRepo)

	// API V1 group
	api := router.PathPrefix("/api/v1").Subrouter()

	// Group: /auth
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", r.AuthHandler.Register).Methods("POST")
	auth.HandleFunc("/login", r.AuthHandler.Login).Methods("POST")

	// Group: /users
	user := api.PathPrefix("/users").Subrouter()
	user.Use(jwtMiddleware)
	user.HandleFunc("", r.UserHandler.ListUsers).Methods("GET")
	user.HandleFunc("", r.UserHandler.CreateUser).Methods("POST")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.GetUserByID).Methods("GET")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.UpdateUser).Methods("PUT")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.DeleteUser).Methods("DELETE")

	// Group: /roles
	role := api.PathPrefix("/roles").Subrouter()
	role.Use(jwtMiddleware)
	role.HandleFunc("", r.RoleHandler.ListRoles).Methods("GET")
	role.HandleFunc("", r.RoleHandler.CreateRole).Methods("POST")
	role.HandleFunc("/{id:[0-9]+}", r.RoleHandler.GetRoleByID).Methods("GET")
	role.HandleFunc("/{id:[0-9]+}", r.RoleHandler.UpdateRole).Methods("PUT")
	role.HandleFunc("/{id:[0-9]+}", r.RoleHandler.DeleteRole).Methods("DELETE")

	// Group: /permissions
	permission := api.PathPrefix("/permissions").Subrouter()
	permission.Use(jwtMiddleware)
	permission.HandleFunc("", r.PermissionHandler.ListPermissions).Methods("GET")
	permission.HandleFunc("", r.PermissionHandler.CreatePermission).Methods("POST")
	permission.HandleFunc("/{id:[0-9]+}", r.PermissionHandler.GetPermissionByID).Methods("GET")
	permission.HandleFunc("/{id:[0-9]+}", r.PermissionHandler.UpdatePermission).Methods("PUT")
	permission.HandleFunc("/{id:[0-9]+}", r.PermissionHandler.DeletePermission).Methods("DELETE")

	// Group: /authorize
	authorize := api.PathPrefix("/authorize").Subrouter()
	authorize.Use(jwtMiddleware)
	authorize.HandleFunc("/assign-role", r.AuthorizeHandler.AssignRole).Methods("POST")
	authorize.HandleFunc("/revoke-role", r.AuthorizeHandler.RevokeRole).Methods("POST")
	authorize.HandleFunc("/assign-permission", r.AuthorizeHandler.AssignPermission).Methods("POST")
	authorize.HandleFunc("/revoke-permission", r.AuthorizeHandler.RevokePermission).Methods("POST")

	return router
}
