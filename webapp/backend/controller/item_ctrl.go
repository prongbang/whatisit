package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type ItemCtrl struct {
	Db *gorm.DB
}

func (ctrl ItemCtrl) GetAll(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) GetByPaged(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) GetLast(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) GetFirst(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) GetById(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) Create(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) CreateOrUpdate(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) Update(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) DeleteAll(c echo.Context) error {
	return nil
}

func (ctrl ItemCtrl) DeleteById(c echo.Context) error {
	return nil
}
