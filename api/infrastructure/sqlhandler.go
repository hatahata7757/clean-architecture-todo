package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"api/interfaces/database"
)

// SqlHandler is a type that includes the gorm.DB
type SqlHandler struct {
	Conn *gorm.DB
}

// NewSqlHandler method connects to the mysql database and returns an mysql instance
func NewSqlHandler() database.SqlHandler {
	conn, err := gorm.Open("mysql", "root:password@tcp(db:3306)/todo")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

// Find method is an instance of Find in the SqlHandler interface
func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// Create method is an instance of Create in the SqlHandler interface
func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}
