package controller

import (
	"net/http"
	"whatisit/webapp/backend/service"
	"whatisit/webapp/backend/utils"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type CategoryCtrl struct {
	Db *gorm.DB
}

func (ctrl CategoryCtrl) GetAll(c echo.Context) error {
	mService := service.CategoryService{ctrl.Db}
	acceptLanguage := c.Request().Header.Get(utils.ACCEPT_LANGUAGE)
	if acceptLanguage != "" {

		return c.JSON(http.StatusOK, mService.GetAllByLangKey(acceptLanguage))
	}
	return c.JSON(http.StatusOK, mService.GetAll())
}

func (ctrl CategoryCtrl) GetByPaged(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) GetLast(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) GetFirst(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) GetById(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) Create(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) CreateOrUpdate(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) Update(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) DeleteAll(c echo.Context) error {
	return nil
}

func (ctrl CategoryCtrl) DeleteById(c echo.Context) error {
	return nil
}
