package routes

import (
	"github.com/hypebeast/goji-boilerplate/controllers"
	"github.com/zenazn/goji"
)

func Include() {
	goji.Get("/", controllers.Home)
}
