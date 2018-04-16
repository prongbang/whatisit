package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type MenuCtrl struct {
	Db *gorm.DB
}

func (ctrl MenuCtrl) GetAll(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) GetByPaged(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) GetLast(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) GetFirst(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) GetById(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) Create(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) CreateOrUpdate(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) Update(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) DeleteAll(c echo.Context) error {
	return nil
}

func (ctrl MenuCtrl) DeleteById(c echo.Context) error {
	return nil
}
