package main

import (
	"whatisit/webapp/backend/controller"
	"whatisit/webapp/backend/db"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	catCtrl := controller.CategoryCtrl{db.Open()}
	mnuCtrl := controller.MenuCtrl{db.Open()}

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
	v1.GET("/menu", mnuCtrl.GetAll) // all, last, first, paged
	v1.GET("/menu/:id", mnuCtrl.GetById)
	v1.POST("/menu", mnuCtrl.Create)
	v1.PUT("/menu/:id", mnuCtrl.CreateOrUpdate)
	v1.PATCH("/menu", mnuCtrl.Update)
	v1.DELETE("/menu", mnuCtrl.DeleteAll)
	v1.DELETE("/menu/:id", mnuCtrl.DeleteById)

	e.Logger.Fatal(e.Start(":1323"))
}
