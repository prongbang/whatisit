package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var session *gorm.DB

func Open() *gorm.DB {
	var err error
	session, err := gorm.Open("mysql", "root:l1ackme@/whatisitdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	return session
}

func Close() {
	session.Close()
}
