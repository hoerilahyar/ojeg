package handler

import (
	"encoding/json"
	"net/http"
	"ojeg/internal/dto"
	"ojeg/internal/usecase"
	"ojeg/pkg/response"
)

type AuthorizeHandler struct {
	Usecase usecase.AuthorizeUsecase
}

func NewAuthorizeHandler(u usecase.AuthorizeUsecase) *AuthorizeHandler {
	return &AuthorizeHandler{Usecase: u}
}

func (h *AuthorizeHandler) AssignRole(w http.ResponseWriter, r *http.Request) {
	var req dto.AssignRoleDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.AssignRole(r.Context(), req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *AuthorizeHandler) RevokeRole(w http.ResponseWriter, r *http.Request) {
	var req dto.RevokeRoleDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.RevokeRole(r.Context(), req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *AuthorizeHandler) AssignPermission(w http.ResponseWriter, r *http.Request) {
	var req dto.AssignPermissionDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.AssignPermission(r.Context(), req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *AuthorizeHandler) RevokePermission(w http.ResponseWriter, r *http.Request) {
	var req dto.RevokePermissionDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.RevokePermission(r.Context(), req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
