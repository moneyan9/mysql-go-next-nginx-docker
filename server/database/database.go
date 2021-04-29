package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	dsn := "user:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	return db
}
