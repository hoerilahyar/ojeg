package registry

import (
	"ojeg/configs"
	"ojeg/delivery/http/handler"
	"ojeg/infrastructure/db"
	"ojeg/infrastructure/jwt"
	"ojeg/internal/service"
	"ojeg/internal/usecase"
	"os"
)

type HandlerRegistry struct {
	UserHandler       *handler.UserHandler
	AuthHandler       *handler.AuthHandler
	RoleHandler       *handler.RoleHandler
	PermissionHandler *handler.PermissionHandler
	AuthorizeHandler  *handler.AuthorizeHandler
}

func NewHandlerRegistry(database *db.DB) *HandlerRegistry {

	secret := os.Getenv("JWT_SECRET")
	issuer := configs.LoadConfig().AppName

	// Repositories
	userRepo := db.NewUserRepository(database)
	roleRepo := db.NewRoleRepository(database)
	permissionRepo := db.NewPermissionRepository(database)
	authorizeRepo := db.NewAuthorizeRepository(database)

	// Services
	userService := service.UserService(userRepo)
	authService := service.AuthService(userRepo, jwt.NewJWTService(secret, issuer))
	roleService := service.RoleService(roleRepo)
	permissionService := service.PermissionService(permissionRepo)
	authorizeService := service.AuthorizeService(authorizeRepo)

	// Usecases
	userUsecase := usecase.UserUsecase(userService)
	authUsecase := usecase.AuthUsecase(authService)
	roleUsecase := usecase.RoleUsecase(roleService)
	permissionUsecase := usecase.PermissionUsecase(permissionService)
	authorizeUsecase := usecase.AuthorizeUsecase(authorizeService)

	// Handlers
	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	roleHandler := handler.NewRoleHandler(roleUsecase)
	permissionHandler := handler.NewPermissionHandler(permissionUsecase)
	authorizeHandler := handler.NewAuthorizeHandler(authorizeUsecase)

	return &HandlerRegistry{
		UserHandler:       userHandler,
		AuthHandler:       authHandler,
		RoleHandler:       roleHandler,
		PermissionHandler: permissionHandler,
		AuthorizeHandler:  authorizeHandler,
	}
}
