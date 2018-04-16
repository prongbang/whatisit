package model

import "github.com/jinzhu/gorm"

type Context struct {
	Db *gorm.DB
}
