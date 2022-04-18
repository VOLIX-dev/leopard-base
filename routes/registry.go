package routes

import "github.com/volix-dev/leopard"

var a *leopard.LeopardApp

func Register(app *leopard.LeopardApp) {
	a = app

	webRoutes()
}
