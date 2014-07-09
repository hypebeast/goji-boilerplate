package routes

import (
	"github.com/hypebeast/goji-boilerplate/apps/bookmarks/controllers"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func Bookmarks(m *web.Mux) {
	goji.Handle("/bookmarks/*", m)
	goji.Get("/bookmarks", http.RedirectHandler("/bookmarks/", 301))

	m.Get("/bookmarks/", controllers.BookmarksHome)
}
