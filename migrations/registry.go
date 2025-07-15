package migrations

import "gorm.io/gorm"

type Migration struct {
	Name string
	Up   func(db *gorm.DB) error
	Down func(db *gorm.DB) error
}

var MigrationList = []Migration{
	{
		Name: "create_users_table",
		Up:   Up_create_users_table,
		Down: Down_create_users_table,
	},
	// Add more migrations here manually

	{
		Name: "role",
		Up:   Up_Create_role_table,
		Down: Down_Create_role_table,
	},

}
