package infrastructure

import (
	"server/infrastructure/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() {
	dsn := "user:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.Migrator().CreateTable(entities.User{})
}
