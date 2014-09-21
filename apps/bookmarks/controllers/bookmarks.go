package controllers

import (
	"net/http"

	"../../../render"
)

func BookmarksHome(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "apps/bookmarks/views/home.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Bookmarks"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
