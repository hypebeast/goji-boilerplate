package routes

import (
	"../apps/bookmarks/routes"
	"../controllers"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Include() {
	goji.Get("/", controllers.Home)
	goji.Get("/about", controllers.About)

	// Bookmarks app
	bookmarks := web.New()
	routes.Bookmarks(bookmarks)
}
