package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func connect(){
	d, err := gorm.Open("mysql", "amrm:password/simplerest?charset=utf8&parseTime=True&loc=local")
	if Err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}