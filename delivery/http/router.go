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
	user.Handle("", jwt.Authorize("user:list", nil)(http.HandlerFunc(r.UserHandler.ListUsers))).Methods("GET")
	user.HandleFunc("", r.UserHandler.CreateUser).Methods("POST")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.GetUserByID).Methods("GET")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.UpdateUser).Methods("PUT")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.DeleteUser).Methods("DELETE")

	// Group: /authorize
	authorize := api.PathPrefix("/authorize").Subrouter()
	authorize.Use(jwtMiddleware)
	authorize.Handle("/assign-role", jwt.Authorize("super-admin:assign-role", nil)(http.HandlerFunc(r.AuthorizeHandler.AssignRole))).Methods("POST")
	authorize.Handle("/revoke-role", jwt.Authorize("super-admin:revoke-role", nil)(http.HandlerFunc(r.AuthorizeHandler.RevokeRole))).Methods("POST")
	authorize.Handle("/assign-permission", jwt.Authorize("super-admin:assign-permission", nil)(http.HandlerFunc(r.AuthorizeHandler.AssignPermission))).Methods("POST")
	authorize.Handle("/revoke-permission", jwt.Authorize("super-admin:revoke-permission", nil)(http.HandlerFunc(r.AuthorizeHandler.RevokePermission))).Methods("POST")

	// Group: /roles
	role := authorize.PathPrefix("/roles").Subrouter()
	role.Use(jwtMiddleware)
	role.Handle("", jwt.Authorize("super-admin:role-list", nil)(http.HandlerFunc(r.RoleHandler.ListRoles))).Methods("GET")
	role.Handle("", jwt.Authorize("super-admin:role-create", nil)(http.HandlerFunc(r.RoleHandler.CreateRole))).Methods("POST")
	role.Handle("/{id:[0-9]+}", jwt.Authorize("super-admin:role-detail", nil)(http.HandlerFunc(r.RoleHandler.GetRoleByID))).Methods("GET")
	role.Handle("/{id:[0-9]+}", jwt.Authorize("super-admin:role-update", nil)(http.HandlerFunc(r.RoleHandler.UpdateRole))).Methods("PUT")
	role.Handle("/{id:[0-9]+}", jwt.Authorize("super-admin:role-delete", nil)(http.HandlerFunc(r.RoleHandler.DeleteRole))).Methods("DELETE")

	// Group: /permissions
	permission := authorize.PathPrefix("/permissions").Subrouter()
	permission.Use(jwtMiddleware)
	permission.Handle("", jwt.Authorize("super-admin:permission-list", nil)(http.HandlerFunc(r.PermissionHandler.ListPermissions))).Methods("GET")
	permission.Handle("", jwt.Authorize("super-admin:permission-create", nil)(http.HandlerFunc(r.PermissionHandler.CreatePermission))).Methods("POST")
	permission.Handle("/{id:[0-9]+}", jwt.Authorize("super-admin:permission-detail", nil)(http.HandlerFunc(r.PermissionHandler.GetPermissionByID))).Methods("GET")
	permission.Handle("/{id:[0-9]+}", jwt.Authorize("super-admin:permission-update", nil)(http.HandlerFunc(r.PermissionHandler.UpdatePermission))).Methods("PUT")
	permission.Handle("/{id:[0-9]+}", jwt.Authorize("super-admin:permission-delete", nil)(http.HandlerFunc(r.PermissionHandler.DeletePermission))).Methods("DELETE")

	return router
}
