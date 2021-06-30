package routes

import (
	controller "Admin-Tools-Api/Controller"
	"github.com/julienschmidt/httprouter"
)

func Customize(app *httprouter.Router) {
	app.GET("/Customize/GetAll", controller.FindAllCustomize)
	app.GET("/Customize/GetById", controller.FindByIdCustomize)
	// app.POST("/Customize/Create", controller.InsertedCustomize)
	// app.PUT("/Customize/Update", controller.UpdatedCustomize)
	// app.DELETE("/Customize/Delete", controller.RemovedCustomize)
}
