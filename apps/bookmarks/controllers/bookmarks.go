package controllers

import (
	"net/http"

	"../../../helpers"
)

func BookmarksHome(w http.ResponseWriter, req *http.Request) {
	templates := helpers.GetBaseTemplates()
	templates = append(templates, "apps/bookmarks/views/home.html")
	err := helpers.RenderTemplate(w, templates, "base", map[string]string{"Title": "Bookmarks"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
