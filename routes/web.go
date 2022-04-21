package routes

import "project_name/controllers"

func webRoutes() {
	a.StaticDir("/assets", "./static")

	a.GET("/", controllers.HomeController, "home")
}
