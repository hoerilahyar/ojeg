package domain

type RolePermission struct {
	ID           uint `gorm:"primaryKey"`
	RoleID       uint
	PermissionID uint
}
