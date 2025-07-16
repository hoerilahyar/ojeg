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

type PermissionHandler struct {
	Usecase usecase.PermissionUsecase
}

func NewPermissionHandler(u usecase.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{Usecase: u}
}

func (h *PermissionHandler) ListPermissions(w http.ResponseWriter, r *http.Request) {
	data, err := h.Usecase.ListPermissions(r.Context())

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func (h *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var req domain.Permission
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}
	err := h.Usecase.CreatePermission(r.Context(), &req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *PermissionHandler) GetPermissionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)
	data, err := h.Usecase.GetPermissionByID(r.Context(), uint(id))

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func (h *PermissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	var req domain.Permission
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, err)
		return
	}

	err := h.Usecase.UpdatePermission(r.Context(), &req)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *PermissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)
	err := h.Usecase.DeletePermission(r.Context(), uint(id))

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
