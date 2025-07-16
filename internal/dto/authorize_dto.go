package dto

type AssignRoleDTO struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type RevokeRoleDTO struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type AssignPermissionDTO struct {
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}

type RevokePermissionDTO struct {
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}
type RoleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PermissionDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
