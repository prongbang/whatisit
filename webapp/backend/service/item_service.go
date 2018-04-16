package service

import (
	"whatisit/webapp/backend/dao"
	"whatisit/webapp/backend/model"

	"github.com/jinzhu/gorm"
)

type ItemService struct {
	Db *gorm.DB
}

func (ctx ItemService) GetAll() []model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.GetAll()
}

func (ctx ItemService) GetAllByLangKey(langKey string) []model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.GetAllByLangKey(langKey)
}

func (ctx ItemService) GetByPaged(limit int, pageSize int) []model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.GetByPaged(limit, pageSize)
}

func (ctx ItemService) GetLast() model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.GetLast()
}

func (ctx ItemService) GetFirst() model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.GetFirst()
}

func (ctx ItemService) GetById(id int64) model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.GetById(id)
}

func (ctx ItemService) Create(item model.Item) model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.Create(item)
}

func (ctx ItemService) CreateOrUpdate(item model.Item) model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.CreateOrUpdate(item)
}

func (ctx ItemService) Update(item model.Item) model.Item {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.Update(item)
}

func (ctx ItemService) DeleteAll() bool {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.DeleteAll()
}

func (ctx ItemService) DeleteById(id int64) bool {
	mDao := dao.ItemDao{ctx.Db}
	return mDao.DeleteById(id)
}
