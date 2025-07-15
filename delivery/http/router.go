package http

import (
	"net/http"

	"ojeg/delivery/http/registry"

	"github.com/gorilla/mux"
)

func NewRouter(r *registry.HandlerRegistry) http.Handler {
	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API V1 group
	api := router.PathPrefix("/api/v1").Subrouter()

	// Group: /users
	user := api.PathPrefix("/users").Subrouter()
	user.HandleFunc("", r.UserHandler.ListUsers).Methods("GET")
	user.HandleFunc("", r.UserHandler.CreateUser).Methods("POST")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.GetUserByID).Methods("GET")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.UpdateUser).Methods("PUT")
	user.HandleFunc("/{id:[0-9]+}", r.UserHandler.DeleteUser).Methods("DELETE")

	// Group: /auth
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", r.AuthHandler.Register).Methods("POST")
	auth.HandleFunc("/login", r.AuthHandler.Login).Methods("POST")

	return router
}
