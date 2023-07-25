package config
//The aim of this file is the handle and return the database and help other files to interact with the database
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:123456789/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}