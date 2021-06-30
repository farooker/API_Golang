package routes

import (
	controller "Admin-Tools-Api/_Controller"
	"github.com/julienschmidt/httprouter"
)

func WS(app *httprouter.Router) {
	app.GET("/websocket", controller.Websocket)
}