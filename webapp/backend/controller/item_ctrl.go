package controller

import (
	"net/http"
	"strconv"
	"whatisit/webapp/backend/model"
	"whatisit/webapp/backend/service"
	"whatisit/webapp/backend/utils"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type ItemCtrl struct {
	Db *gorm.DB
}

func (ctrl ItemCtrl) GetAll(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	acceptLanguage := c.Request().Header.Get(utils.ACCEPT_LANGUAGE)
	if acceptLanguage != "" {

		return c.JSON(http.StatusOK, mService.GetAllByLangKey(acceptLanguage))
	}
	return c.JSON(http.StatusOK, mService.GetAll())
}

func (ctrl ItemCtrl) GetByPaged(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	acceptLanguage := c.Request().Header.Get(utils.ACCEPT_LANGUAGE)
	if acceptLanguage != "" {
		limit := c.Param("limit")
		pageSize := c.Param("pageSize")
		if limit != "" && pageSize != "" {
			limitInt, err1 := strconv.Atoi(limit)
			pageSizeInt, err2 := strconv.Atoi(pageSize)
			if err1 == nil && err2 == nil {
				return c.JSON(http.StatusOK, mService.GetByPaged(limitInt, pageSizeInt))
			}
		}
	}
	return c.JSON(http.StatusOK, mService.GetAll())
}

func (ctrl ItemCtrl) GetLast(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	return c.JSON(http.StatusOK, mService.GetLast())
}

func (ctrl ItemCtrl) GetFirst(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	return c.JSON(http.StatusOK, mService.GetFirst())
}

func (ctrl ItemCtrl) GetById(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	if id := c.Param("id"); id != "" {
		if id64, err := strconv.ParseInt(id, 10, 64); err == nil && id64 != 0 {
			return c.JSON(http.StatusOK, mService.GetById(id64))
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Can't get item by " + id})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Can't get item"})
}

func (ctrl ItemCtrl) Create(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	var item model.Item
	if err := c.Bind(&item); err == nil {
		return c.JSON(http.StatusOK, mService.Create(item))
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Can't create item"})
}

func (ctrl ItemCtrl) CreateOrUpdate(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	var item model.Item
	id := c.Param("id")
	if err := c.Bind(&item); err == nil {
		if id64, err := strconv.ParseInt(id, 10, 64); err == nil && id64 != 0 {
			item.ID = id64
		}
		return c.JSON(http.StatusOK, mService.CreateOrUpdate(item))
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Can't update item"})
}

func (ctrl ItemCtrl) Update(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	var item model.Item
	id := c.Param("id")
	if err := c.Bind(&item); err == nil && id != "" {
		if id64, err := strconv.ParseInt(id, 10, 64); err == nil && id64 != 0 {
			item.ID = id64
			return c.JSON(http.StatusOK, mService.Update(item))
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Can't update item"})
}

func (ctrl ItemCtrl) DeleteAll(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	if mService.DeleteAll() {
		return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Fail"})
}

func (ctrl ItemCtrl) DeleteById(c echo.Context) error {
	mService := service.ItemService{ctrl.Db}
	id := c.Param("id")
	errorMessage := "ID Not found."
	if id != "" {
		id64, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			if mService.DeleteById(id64) {
				return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
			}
			return c.JSON(http.StatusOK, map[string]string{"message": "Fail"})
		}
		errorMessage = err.Error()
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"message": errorMessage})
}
