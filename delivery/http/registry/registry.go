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
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
}

func NewHandlerRegistry(database *db.DB) *HandlerRegistry {

	secret := os.Getenv("JWT_SECRET")
	issuer := configs.LoadConfig().AppName

	// Repositories
	userRepo := db.NewUserRepository(database)

	// Services
	userService := service.UserService(userRepo)
	authService := service.AuthService(userRepo, jwt.NewJWTService(secret, issuer))

	// Usecases
	userUsecase := usecase.UserUsecase(userService)
	authUsecase := usecase.AuthUsecase(authService)

	// Handlers
	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)

	return &HandlerRegistry{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
