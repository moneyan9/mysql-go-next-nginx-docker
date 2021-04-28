package helpers

import (
	"gorm.io/driver/mysql"
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
	conn *gorm.DB
}

func NewSqlHandler() SqlHandler {
	dsn := "user:password@tcp(db:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	return &sqlHandler{conn: conn}
}

func (handler *sqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.conn.Find(out, where...)
}

func (handler *sqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.conn.Exec(sql, values...)
}

func (handler *sqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.conn.First(out, where...)
}

func (handler *sqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.conn.Raw(sql, values...)
}

func (handler *sqlHandler) Create(value interface{}) *gorm.DB {
	return handler.conn.Create(value)
}

func (handler *sqlHandler) Save(value interface{}) *gorm.DB {
	return handler.conn.Save(value)
}

func (handler *sqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.conn.Delete(value)
}

func (handler *sqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.conn.Where(query, args...)
}
