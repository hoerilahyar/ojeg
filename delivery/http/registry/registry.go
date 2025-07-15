package registry

import (
	"ojeg/delivery/http/handler"
	"ojeg/infrastructure/db"
	"ojeg/internal/user/service"
)

type HandlerRegistry struct {
	UserHandler *handler.UserHandler
	// Tambahkan handler lain nanti: DriverHandler, BookingHandler, etc
}

func NewHandlerRegistry(database *db.DB) *HandlerRegistry {
	// User module wiring
	userRepo := db.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &HandlerRegistry{
		UserHandler: userHandler,
	}
}
