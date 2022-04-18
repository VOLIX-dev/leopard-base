package main

import (
	"github.com/volix-dev/leopard"
	"project_name/routes"
)

func main() {
	app, err := leopard.New()

	if err != nil {
		panic(err)
	}

	routes.Register(app)

	err = app.Serve()

	if err != nil {
		panic(err)
	}
}
