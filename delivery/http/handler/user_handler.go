package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ojeg/internal/domain"
	"ojeg/internal/usecase"
	"ojeg/pkg/errors"
	"ojeg/pkg/response"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.ListUsers(r.Context())
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, users)
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse input JSON to User struct
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	// Call the correct method
	err := h.usecase.CreateUser(r.Context(), &user)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, user)
}

// GetUserByID handles GET /users/{id}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, errors.ErrValNotString)
		return
	}

	user, err := h.usecase.GetUserByID(r.Context(), uint(id))
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, user)
}

// UpdateUser handles PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, errors.ErrValNotString)
		return
	}

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		response.Error(w, errors.ErrInvalidPayload)
		return
	}
	user.ID = uint(id)

	err = h.usecase.UpdateUser(r.Context(), &user)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, user)
}

// DeleteUser handles DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	if err := h.usecase.DeleteUser(r.Context(), uint(id)); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
