package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ojeg/internal/domain"
	"ojeg/internal/usecase"
	"ojeg/pkg/response"

	"github.com/gorilla/mux"
)

type RoleHandler struct {
	Usecase usecase.RoleUsecase
}

func NewRoleHandler(u usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{Usecase: u}
}

func (h *RoleHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	data, err := h.Usecase.ListRoles(r.Context())

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req domain.Role
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.CreateRole(r.Context(), &req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *RoleHandler) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)
	data, err := h.Usecase.GetRoleByID(r.Context(), uint(id))

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	var req domain.Role
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.UpdateRole(r.Context(), &req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)
	err := h.Usecase.DeleteRole(r.Context(), uint(id))

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
