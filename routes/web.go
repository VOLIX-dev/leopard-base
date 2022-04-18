package routes

import "project_name/controllers"

func webRoutes() {
	a.StaticDir("/static", "./static")

	a.GET("/", controllers.HomeController, "home")
}
