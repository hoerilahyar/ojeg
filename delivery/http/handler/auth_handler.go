package handler

import (
	"encoding/json"
	"net/http"

	"ojeg/internal/domain"
	"ojeg/internal/usecase"
	"ojeg/pkg/errors"
	"ojeg/pkg/response"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	err := h.usecase.Register(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, errors.ErrInvalidPayload)
		return
	}

	user, err := h.usecase.Login(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, user)
}
