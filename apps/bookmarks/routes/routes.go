package routes

import (
	"../controllers"

	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Bookmarks(m *web.Mux) {
	goji.Handle("/bookmarks/*", m)
	goji.Get("/bookmarks", http.RedirectHandler("/bookmarks/", 301))

	m.Get("/bookmarks/", controllers.BookmarksHome)
}
