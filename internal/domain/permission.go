package domain

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(191);uniqueIndex;not null"`
}
