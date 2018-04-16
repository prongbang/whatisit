package service

import (
	"whatisit/webapp/backend/dao"
	"whatisit/webapp/backend/model"

	"github.com/jinzhu/gorm"
)

type CategoryService struct {
	Db *gorm.DB
}

func (ctx CategoryService) GetAll() []model.Category {
	mDao := dao.CatagoryDao{ctx.Db}
	return mDao.GetAll()
}

func (ctx CategoryService) GetAllByLangKey(langKey string) []model.Category {
	mDao := dao.CatagoryDao{ctx.Db}

	return mDao.GetAllByLangKey(langKey)
}
