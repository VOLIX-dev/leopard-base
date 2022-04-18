package controllers

import "github.com/volix-dev/leopard"

func HomeController(c leopard.ContextInterface) {
	err := c.RenderTemplate("home.twig", nil)

	if err != nil {
		_ = c.Error(err)
		return
	}
}
