package configs

import (
	"server/database"
	"server/entities"
)

func InitializeDatabase() {
	migrator := database.GetDatabase().Migrator()
	migrator.CreateTable(entities.User{})
}
