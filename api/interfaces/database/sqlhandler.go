package database

import "github.com/jinzhu/gorm"

// SqlHandler is an interface that defines the behavior of gorm.DB
type SqlHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
}
