package routes

import (
	controller "Admin-Tools-Api/_Controller"
	"github.com/julienschmidt/httprouter"
)

func FileAccess(app *httprouter.Router) {
	app.GET("/FileAccessWrite", controller.WriteFileAccess)
	app.GET("/FileAccessRead", controller.ReadFileAccess)
}