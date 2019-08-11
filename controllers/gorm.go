package controllers

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

type InDB struct {
	DB *gorm.DB
}
