package dao

import (
	"whatisit/webapp/backend/model"

	"github.com/jinzhu/gorm"
)

type CatagoryDao struct {
	Db *gorm.DB
}

func (ctx CatagoryDao) GetAll() []model.Category {

	var category []model.Category
	ctx.Db.Find(&category)

	return category
}

func (ctx CatagoryDao) GetAllByLangKey(langKey string) []model.Category {

	var category []model.Category
	ctx.Db.Where("lang_key = ?", langKey).Find(&category)

	return category
}
