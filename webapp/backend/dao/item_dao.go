package dao

import (
	"whatisit/webapp/backend/model"

	"github.com/jinzhu/gorm"
)

type ItemDao struct {
	Db *gorm.DB
}

func (ctx ItemDao) GetAll() []model.Item {

	var item []model.Item
	ctx.Db.Order("id desc").Find(&item)
	return item
}

func (ctx ItemDao) GetAllByLangKey(langKey string) []model.Item {

	var item []model.Item
	ctx.Db.Where("lang_key = ?", langKey).Order("id desc").Find(&item)
	return item
}

func (ctx ItemDao) GetByPaged(limit int, pageSize int) []model.Item {

	var item []model.Item
	ctx.Db.Limit(limit).Offset(pageSize).Find(&item)
	return item
}

func (ctx ItemDao) GetLast() model.Item {

	var item model.Item
	ctx.Db.Last(&item)
	return item
}

func (ctx ItemDao) GetFirst() model.Item {

	var item model.Item
	ctx.Db.First(&item)
	return item
}

func (ctx ItemDao) GetById(id int64) model.Item {

	var item model.Item
	ctx.Db.Where("id = ?", id).Find(&item)
	return item
}

func (ctx ItemDao) Create(item model.Item) model.Item {

	ctx.Db.Create(item)
	return item
}

func (ctx ItemDao) CreateOrUpdate(item model.Item) model.Item {

	data := ctx.GetById(item.ID)
	if data.ID != 0 {
		ctx.Update(item)
		return item
	}
	ctx.Create(item)
	return item
}

func (ctx ItemDao) Update(item model.Item) model.Item {

	ctx.Db.Update(item)
	return item
}

func (ctx ItemDao) DeleteAll() bool {

	ctx.Db.Delete(model.Item{})
	return true
}

func (ctx ItemDao) DeleteById(id int64) bool {

	ctx.Db.Where("id = ?", id).Delete(model.Item{})
	return true
}
