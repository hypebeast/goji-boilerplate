package controllers

import (
	"../render"

	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "views/index.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func About(w http.ResponseWriter, req *http.Request) {
	// TODO
}
