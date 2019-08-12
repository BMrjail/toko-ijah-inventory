package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBInit create connection to database
func DBInit() *gorm.DB {

		db, err := gorm.Open("sqlite3", "toko_ijah.db")

		//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/pet?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect to database")
		}

		return db


	//db.AutoMigrate(structs.Person{})
	//db.AutoMigrate(structs.Berita{})

}
