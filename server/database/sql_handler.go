package database

import (
	"gorm.io/gorm"
)

type SqlHandler interface {
	Exec(string, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
}

type sqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler(db *gorm.DB) SqlHandler {
	return &sqlHandler{db: db}
}

func (handler *sqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(out, where...)
}

func (handler *sqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.db.Exec(sql, values...)
}

func (handler *sqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.First(out, where...)
}

func (handler *sqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.db.Raw(sql, values...)
}

func (handler *sqlHandler) Create(value interface{}) *gorm.DB {
	return handler.db.Create(value)
}

func (handler *sqlHandler) Save(value interface{}) *gorm.DB {
	return handler.db.Save(value)
}

func (handler *sqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.db.Delete(value)
}

func (handler *sqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.db.Where(query, args...)
}
