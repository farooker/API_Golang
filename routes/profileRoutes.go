package routes

import (
	controller "Admin-Tools-Api/controller"

	"github.com/julienschmidt/httprouter"
)

func Profiles(app *httprouter.Router) {
	app.GET("/GetAll", controller.FindProfiles)
	app.GET("/GetById", controller.FindByIdProfiles)
	app.GET("/Create", controller.InsertedProfiles)
	app.GET("/Update", controller.UpdatedProfiles)
	app.GET("/Delete", controller.RemovedProfiles)
}
