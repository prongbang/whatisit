package main

import (
	"whatisit/webapp/backend/controller"
	"whatisit/webapp/backend/db"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := db.Open()
	catCtrl := controller.CategoryCtrl{db}
	itmCtrl := controller.ItemCtrl{db}

	v1 := e.Group("/api/v1")
	// Category
	v1.GET("/category", catCtrl.GetAll) // all, last, first, paged
	v1.GET("/category/:id", catCtrl.GetById)
	v1.POST("/category", catCtrl.Create)
	v1.PUT("/category/:id", catCtrl.CreateOrUpdate)
	v1.PATCH("/category", catCtrl.Update)
	v1.DELETE("/category", catCtrl.DeleteAll)
	v1.DELETE("/category/:id", catCtrl.DeleteById)

	// Menu
	v1.GET("/menu", itmCtrl.GetAll) // all, last, first, paged
	v1.GET("/menu/:id", itmCtrl.GetById)
	v1.POST("/menu", itmCtrl.Create)
	v1.PUT("/menu/:id", itmCtrl.CreateOrUpdate)
	v1.PATCH("/menu", itmCtrl.Update)
	v1.DELETE("/menu", itmCtrl.DeleteAll)
	v1.DELETE("/menu/:id", itmCtrl.DeleteById)

	e.Logger.Fatal(e.Start(":1323"))
}
