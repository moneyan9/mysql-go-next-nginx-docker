package database

import (
	"server/entities"
)

func InitializeDatabase() {
	migrator := GetDatabase().Migrator()
	migrator.CreateTable(
		entities.Group{},
		entities.User{},
	)
}
